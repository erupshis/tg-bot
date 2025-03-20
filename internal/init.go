package internal

import (
	"fmt"
	"net/http"
	"os"

	"github.com/erupshis/tg-bot/internal/bot"
	"github.com/sirupsen/logrus"
)

func (a *App) Init() error {
	if err := initLogrus(a.cfg.LogLevel); err != nil {
		return fmt.Errorf("init logger: %w", err)
	}

	var err error
	a.tgBot, err = bot.NewTelegramBot(a.cfg.BotToken, a.cfg.YCID, true)
	if err != nil {
		return fmt.Errorf("create telegram bot: %w", err)
	}

	a.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", a.cfg.YCPort),
		Handler: nil,
	}

	return nil
}

func initLogrus(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("parse Level from config: %w", err)
	}

	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	return nil
}
