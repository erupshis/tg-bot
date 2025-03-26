package main

import (
	"context"

	"github.com/erupshis/tg-bot/internal"
	"github.com/erupshis/tg-bot/internal/pkg/closer"
)

func main() {
	ctxWithCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := internal.NewApp()
	// graceShutdown.
	closer.Add(func() { _ = app.Shutdown(ctxWithCancel) })

	app.Run(ctxWithCancel)
}
