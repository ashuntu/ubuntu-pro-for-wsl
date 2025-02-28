// Package registrywatcher implements a service that updates the config every time the registry changes.
package registrywatcher

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/canonical/ubuntu-pro-for-wsl/windows-agent/internal/config"
	"github.com/canonical/ubuntu-pro-for-wsl/windows-agent/internal/distros/database"
	log "github.com/canonical/ubuntu-pro-for-wsl/windows-agent/internal/grpc/logstreamer"
	"github.com/canonical/ubuntu-pro-for-wsl/windows-agent/internal/proservices/registrywatcher/registry"
)

// Service is a service that monitors the Windows registry for any changes to the key
// Software/Canonical/UbuntuPro.
//
// If a change is detected, the new contents of the registry key are pushed to the
// config.
type Service struct {
	ctx  context.Context
	stop func()

	running chan struct{}

	registry Registry
	conf     Config
	db       *database.DistroDB
}

var registryPath = []string{"Software", "Canonical", "UbuntuPro"}

// Registry is an interface to the Windows registry.
type Registry interface {
	HKCUOpenKey(path string) (registry.Key, error)
	CloseKey(k registry.Key)
	ReadValue(k registry.Key, field string) (value string, err error)

	// Win32 stuff: not strictly registry but not worth separating out
	RegNotifyChangeKeyValue(k registry.Key) (registry.Event, error)
	WaitForSingleObject(event registry.Event) error
	CloseEvent(ev registry.Event)
}

// Config is an interface to easily allow dependency injection. Should be a config.Config
// in production.
type Config interface {
	UpdateRegistryData(context.Context, config.RegistryData, *database.DistroDB) error
}

type options struct {
	registry Registry
}

// Option is an optional argument for the registry watcher.
type Option = func(*options)

// WithRegistry allows for overriding the registry back-end.
func WithRegistry(r Registry) Option {
	return func(o *options) {
		o.registry = r
	}
}

// New creates a registry watcher service.
func New(ctx context.Context, conf Config, database *database.DistroDB, args ...Option) Service {
	var opts options

	for _, f := range args {
		f(&opts)
	}

	if opts.registry == nil {
		opts.registry = registry.Windows{}
	}

	return Service{
		registry: opts.registry,
		conf:     conf,
		db:       database,

		ctx:     ctx,
		stop:    func() {},
		running: make(chan struct{}),
	}
}

// Start starts watching the service. It does a first read of the registry
// before returning.
func (s *Service) Start() {
	s.ctx, s.stop = context.WithCancel(s.ctx)

	s.readThenPushRegistryData(s.ctx)

	go s.run()
}

// Stop releases all resources associated with the registry watcher.
func (s *Service) Stop() {
	s.stop()
	<-s.running
}

// run is the blocking registry watcher.
func (s *Service) run() {
	defer close(s.running)
	/*
		When we detect a change we don't immediately read the registry and push
		the new data. Instead, we wait until we're watching again. This way we
		avoid silent changes in between ending and starting successive watches.

		In the case we fail to watch, we still push changes just in case. False
		positives don't matter much because the config will ignore data that are
		not new.
	*/

	// These rates are NOT how often we look at the registry. Registry updates are
	// detected instantaneously. Rather, they are to avoid entering a hot loop if
	// we fail to start watching the registry for whatever reason.
	const (
		minRate      = time.Second
		growthFactor = 2
		maxRate      = 30 * time.Minute
	)
	retryRate := minRate

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
		}

		// Returns error if we need to sleep in order to avoid a hot loop.
		err := func() error {
			ctx, cancel := context.WithCancel(s.ctx)
			defer cancel()

			k, path, err := s.openDeepestExistingKey(registryPath)
			if err != nil {
				return err
			}
			defer s.registry.CloseKey(k)

			// Start to watch
			event, err := s.registry.RegNotifyChangeKeyValue(k)
			if err != nil {
				return fmt.Errorf(`could not watch changes to registry key HKCU\%s: %v`, path, err)
			}
			defer s.registry.CloseEvent(event)

			log.Debugf(ctx, `Registry watcher: watching key HKCU\%s`, path)

			// Push update right after having started to watch
			s.readThenPushRegistryData(ctx)

			// Wait until the key is modified or the context is cancelled, whichever one happens first
			if err := s.waitForSingleObject(ctx, event); err != nil {
				return fmt.Errorf(`could not wait for changes to registry key HKCU\%s: %v`, path, err)
			}
			log.Infof(ctx, `Registry watcher: detected change in registry key HKCU\%s or one of its children`, path)

			return nil
		}()

		if err != nil {
			log.Warningf(s.ctx, "Registry watcher: %v", err)
			s.readThenPushRegistryData(s.ctx)

			select {
			case <-s.ctx.Done():
				return
			case <-time.After(retryRate):
			}

			retryRate = min(growthFactor*retryRate, maxRate)
			continue
		}

		retryRate = minRate
	}
}

// openDeepestExistingKey opens the key that is deepest existing key in the tree
// defined by the path.
//
// This intends to solve the problem of 'what if the key does not exist?'.
// AKA: any default install. The solution is to watch its parent directory,
// and wait until our key is created.
func (s *Service) openDeepestExistingKey(registryPath []string) (k registry.Key, path string, err error) {
	for depth := len(registryPath); depth >= 1; depth-- {
		//                          ^~~~~~~~~~
		// We go up to depth 1. Going to depth 0 would watch the entire
		// registry hive and that is overkill.
		path = filepath.Join(registryPath[:depth]...)

		k, err = s.registry.HKCUOpenKey(path)
		if err != nil {
			if errors.Is(err, registry.ErrKeyNotExist) {
				continue
			}
			break
		}
		return k, path, nil
	}

	// We failed to open a key, or we could not find any key along the path
	return 0, "", fmt.Errorf(`could not open registry key HKCU\%s: %v`, path, err)
}

// waitForSingleObject is a utility wrapper around Win32's WaitForSingleObject. It allows
// cancelling the wait with the use of a context.
//
// Cancelling the context skips the wait, but does not release resources. These are released
// once the event is set.
func (s *Service) waitForSingleObject(ctx context.Context, event registry.Event) error {
	ch := make(chan error, 1)

	go func() {
		ch <- s.registry.WaitForSingleObject(event)
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-ch:
		return err
	}
}

// readThenPushRegistryData reads the registry and pushes the read data to the config.
// This function is syntax sugar for Start, so we log the errors instead of having
// the caller deal with them.
func (s *Service) readThenPushRegistryData(ctx context.Context) {
	data, err := loadRegistry(s.registry)
	if err != nil {
		log.Warningf(ctx, "Registry watcher: could not read registry: %v", err)
		return
	}

	if err := s.conf.UpdateRegistryData(ctx, data, s.db); err != nil {
		log.Warningf(ctx, "Registry watcher: could not push new registry data: %v", err)
	}
}

func loadRegistry(reg Registry) (data config.RegistryData, err error) {
	k, err := reg.HKCUOpenKey(filepath.Join(registryPath...))
	if errors.Is(err, registry.ErrKeyNotExist) {
		// Default values
		return data, nil
	}
	if err != nil {
		return data, err
	}
	defer reg.CloseKey(k)

	proToken, err := readFromRegistry(reg, k, "UbuntuProToken")
	if err != nil {
		return data, err
	}

	conf, err := readFromRegistry(reg, k, "LandscapeConfig")
	if err != nil {
		return data, err
	}

	return config.RegistryData{
		UbuntuProToken:  proToken,
		LandscapeConfig: conf,
	}, nil
}

func readFromRegistry(r Registry, key registry.Key, field string) (string, error) {
	value, err := r.ReadValue(key, field)
	if errors.Is(err, registry.ErrFieldNotExist) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("could not read %q from registry", field)
	}

	return value, nil
}
