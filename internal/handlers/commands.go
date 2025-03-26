package handlers

import (
	"fmt"

	"github.com/erupshis/tg-bot/locales"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (m *Manager) StartCommand(bot *tgbotapi.BotAPI, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, m.locales.Get(locales.Messages.Commands.Start))
	if _, err := bot.Send(msg); err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}

func (m *Manager) HelpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, m.locales.Get(locales.Messages.Commands.Help))
	if _, err := bot.Send(msg); err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}
