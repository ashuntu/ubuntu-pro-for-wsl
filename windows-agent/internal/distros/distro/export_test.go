package distro

import (
	"context"

	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/distros/initialTasks"
)

func WithTaskProcessingContext(ctx context.Context) Option {
	return func(o *options) {
		if ctx != nil {
			o.taskProcessingContext = ctx
		}
	}
}

// WithNewWorker is an optional parameter for distro.New that allows for overriding
// the worker.New constructor. It is meant for dependency injection.
func WithNewWorker(newWorkerFunc func(context.Context, *Distro, string, *initialTasks.InitialTasks) (Worker, error)) Option {
	return func(o *options) {
		o.newWorkerFunc = newWorkerFunc
	}
}

// Identity contains persistent and uniquely identifying information about the distro.
type Identity = identity

// GetIdentity returns a reference to the distro's identity.
//
//nolint: revive
// False positive, Identity is exported.
func (d *Distro) GetIdentity() *Identity {
	return &d.identity
}
