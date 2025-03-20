package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/erupshis/tg-bot/internal"
	"github.com/erupshis/tg-bot/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("load envs: %s", err.Error())
	}

	app := internal.NewApp(cfg)

	if err = app.Init(); err != nil {
		log.Fatalf("init app: %s", err.Error())
	}

	ctxWithCancel, cancel := context.WithCancel(context.Background())
	app.Run(ctxWithCancel)

	// graceShutdown.
	graceShutdown(cancel, app)
}

func graceShutdown(cancel context.CancelFunc, app *internal.App) {
	idleConnsClosed := make(chan struct{})
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigCh

		if err := app.Shutdown(context.Background()); err != nil {
			logrus.Errorf("http server graceShutdown: %s", err.Error())
		}

		cancel()
		close(idleConnsClosed)
	}()

	<-idleConnsClosed
	logrus.Infof("service graceShutdown gracefully")
}
