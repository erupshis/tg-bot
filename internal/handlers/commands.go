package handlers

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (m *Manager) StartCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, m.cfg.Greetings)
	if _, err := bot.Send(msg); err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}

func (m *Manager) HelpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Вот что я умею:\n\n"+
		"/start - Начать работу с ботом\n"+
		"/help - Получить справку\n\n"+
		"Просто отправь мне сообщение, и я передам его администратору на проверку.")
	if _, err := bot.Send(msg); err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}
