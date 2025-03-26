package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/erupshis/tg-bot/internal/config"
	"github.com/erupshis/tg-bot/internal/localization"
	"github.com/erupshis/tg-bot/internal/tg_bot"
	"github.com/sirupsen/logrus"
)

func (a *App) InitConfig() *App {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("load envs: %s", err.Error())
	}

	a.cfg = cfg
	return a
}

func (a *App) InitLocales() *App {
	var err error
	a.locales, err = localization.New(a.cfg.Lang)
	if err != nil {
		log.Fatalf("load localization: %s", err.Error())
	}

	return a
}

func (a *App) InitLogger() *App {
	level, err := logrus.ParseLevel(a.cfg.LogLevel)
	if err != nil {
		log.Fatalf("parse log level from config: %s", err.Error())
	}

	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	return a
}

func (a *App) InitTelegramBot() *App {
	var err error
	a.tgBot, err = tg_bot.NewTelegramBot(a.cfg.BotToken, a.cfg.YCID, a.cfg.Debug)
	if err != nil {
		log.Fatalf("create telegram tg_bot: %s", err.Error())
	}

	return a
}

func (a *App) InitHttpServer() *App {
	a.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", a.cfg.YCPort),
		Handler: nil,
	}
	return a
}
