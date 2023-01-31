// Package main is the windows-agent entry point.
package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/cmd/ubuntu-pro-agent/agent"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/consts"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/i18n"
	log "github.com/sirupsen/logrus"
)

//go:generate go run ../generate_completion_documentation.go completion ../../generated
//go:generate go run ../generate_completion_documentation.go update-readme
//go:generate go run ../generate_completion_documentation.go update-doc-cli-ref

func main() {
	a := agent.New()
	os.Exit(run(a))
}

type app interface {
	Run() error
	UsageError() bool
	Quit()
}

func run(a app) int {
	i18n.InitI18nDomain(consts.TEXTDOMAIN)
	defer installSignalHandler(a)()

	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		DisableTimestamp:       true,
	})

	if err := a.Run(); err != nil {
		log.Error(err)

		if a.UsageError() {
			return 2
		}
		return 1
	}

	return 0
}

func installSignalHandler(a app) func() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			switch v, ok := <-c; v {
			case syscall.SIGINT, syscall.SIGTERM:
				a.Quit()
				return
			default:
				// channel was closed: we exited
				if !ok {
					return
				}
			}
		}
	}()

	return func() {
		signal.Stop(c)
		close(c)
		wg.Wait()
	}
}
