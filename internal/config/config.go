package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	YCID      string
	YCPort    string
	BotToken  string
	ChannelID string
	AdminID   int64

	Greetings string

	Debug bool
}

func New() (*Config, error) {
	cfg := Config{}
	cfg.YCID = os.Getenv("YC_ID")
	if cfg.YCID == "" {
		return nil, fmt.Errorf("YC_ID not set")
	}

	cfg.YCPort = os.Getenv("PORT")
	if cfg.YCPort == "" {
		return nil, fmt.Errorf("PORT not set")
	}

	cfg.BotToken = os.Getenv("BOT_TOKEN")
	if cfg.BotToken == "" {
		return nil, fmt.Errorf("bot token not found")
	}

	cfg.ChannelID = os.Getenv("CHANNEL_ID")
	if cfg.ChannelID == "" {
		return nil, fmt.Errorf("channel ID not found")
	}

	cfg.Greetings = os.Getenv("GREETINGS")

	adminIDStr := os.Getenv("ADMIN_ID")
	var err error
	cfg.AdminID, err = strconv.ParseInt(adminIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse AdminID: %w", err)
	}

	return &cfg, nil
}
