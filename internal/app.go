package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/erupshis/tg-bot/internal/bot"
	"github.com/erupshis/tg-bot/internal/config"
	"github.com/erupshis/tg-bot/internal/handlers"
	"github.com/sirupsen/logrus"
)

type App struct {
	tgBot  *bot.Telegram
	server *http.Server

	cfg *config.Config
}

func NewApp(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}

}

func (a *App) Run(ctx context.Context) {
	go func() {
		if err := a.tgBot.Run(ctx, handlers.NewManager(a.cfg)); err != nil {
			if !errors.Is(err, context.Canceled) {
				logrus.Errorf("Telegram bot failed: %s", err.Error())
			}
		}
	}()

	go func() {
		if errServer := a.server.ListenAndServe(); errServer != nil {
			logrus.Errorf("http server serve: %s", errServer.Error())
		}
	}()
}

func (a *App) Shutdown(ctx context.Context) error {
	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("http server shutdown: %w", err)
	}

	return nil
}
