// Package daemon is handling the TCP connection and connecting a GRPC service to it.
package daemon

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/canonical/ubuntu-pro-for-wsl/common"
	"github.com/canonical/ubuntu-pro-for-wsl/common/i18n"
	log "github.com/canonical/ubuntu-pro-for-wsl/windows-agent/internal/grpc/logstreamer"
	"github.com/ubuntu/decorate"
	"google.golang.org/grpc"
)

// GRPCServiceRegisterer is a function that the daemon will call everytime we want to build a new GRPC object.
type GRPCServiceRegisterer func(ctx context.Context) *grpc.Server

// Daemon is a daemon for windows agents with grpc support.
type Daemon struct {
	listeningPortFilePath string

	grpcServer *grpc.Server
}

// New returns an new, initialized daemon server that is ready to register GRPC services.
// It hooks up to windows service management handler.
func New(ctx context.Context, registerGRPCServices GRPCServiceRegisterer, addrDir string) *Daemon {
	log.Debug(ctx, "Building new daemon")

	listeningPortFilePath := filepath.Join(addrDir, common.ListeningPortFileName)
	log.Debugf(ctx, "Daemon port file path: %s", listeningPortFilePath)

	return &Daemon{
		listeningPortFilePath: listeningPortFilePath,
		grpcServer:            registerGRPCServices(ctx),
	}
}

// Serve listens on a tcp socket and starts serving GRPC requests on it.
// Before serving, it writes a file on disk on which port it's listening on for client
// to be able to reach our server.
// This file is removed once the server stops listening.
func (d Daemon) Serve(ctx context.Context) (err error) {
	defer decorate.OnError(&err, i18n.G("error while serving"))

	log.Debug(ctx, "Starting to serve requests")

	// TODO: get a local port only, please :)
	var cfg net.ListenConfig
	lis, err := cfg.Listen(ctx, "tcp", "")
	if err != nil {
		return fmt.Errorf("can't listen: %v", err)
	}

	addr := lis.Addr().String()

	// Write a file on disk to signal selected ports to clients.
	// We write it here to signal error when calling service.Start().
	if err := os.WriteFile(d.listeningPortFilePath, []byte(addr), 0600); err != nil {
		return err
	}
	defer os.Remove(d.listeningPortFilePath)

	log.Infof(ctx, "Serving GRPC requests on %v", addr)

	if err := d.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("grpc error: %v", err)
	}
	return nil
}

// Quit gracefully quits listening loop and stops the grpc server.
// It can drop any existing connexion if force is true.
func (d Daemon) Quit(ctx context.Context, force bool) {
	log.Info(ctx, "Stopping daemon requested.")
	if force {
		d.grpcServer.Stop()
		return
	}

	log.Info(ctx, i18n.G("Wait for active requests to close."))
	d.grpcServer.GracefulStop()
	log.Debug(ctx, i18n.G("All connections have now ended."))
}
