package handlers

import (
	"github.com/erupshis/tg-bot/internal/config"
	"github.com/erupshis/tg-bot/internal/localization"
)

type Manager struct {
	cfg     *config.Config
	locales *localization.Localizer
}

func NewManager(cfg *config.Config, locales *localization.Localizer) *Manager {
	return &Manager{
		cfg:     cfg,
		locales: locales,
	}
}
