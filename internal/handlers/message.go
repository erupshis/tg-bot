package handlers

import (
	"fmt"

	"github.com/erupshis/tg-bot/internal/pkg/text_formatter"
	"github.com/erupshis/tg-bot/locales"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Message Обработка текстовых сообщений
func (m *Manager) Message(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	userMessage := text_formatter.EscapeMarkdownV2(message.Text)

	if !isMessageValid(userMessage, int(m.cfg.MinMessageLen)) {
		response := m.locales.Getf(locales.Messages.User.MessageTooShort, m.cfg.MinMessageLen)
		if _, err := bot.Send(tgbotapi.NewMessage(message.Chat.ID, response)); err != nil {
			return fmt.Errorf("sending confirmation message to user: %w", err)
		}

		return nil
	}

	if err := m.sendMessageToAdmin(bot, userMessage); err != nil {
		return err
	}

	msgChannel := tgbotapi.NewMessageToChannel(m.cfg.ChannelID, userMessage)
	msgChannel.ParseMode = tgbotapi.ModeMarkdownV2
	if _, err := bot.Send(msgChannel); err != nil {
		return fmt.Errorf("sending unchecked message in channel: %w", err)
	}

	// Отправляем подтверждение пользователю
	if _, err := bot.Send(tgbotapi.NewMessage(message.Chat.ID, m.locales.Get(locales.Messages.User.MessageReceived))); err != nil {
		return fmt.Errorf("sending confirmation message to user: %w", err)
	}

	return nil
}

func (m *Manager) sendMessageToAdmin(bot *tgbotapi.BotAPI, userMessage string) error {
	// Создаем inline клавиатуру с кнопками ДА/НЕТ
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(m.locales.Get(locales.Messages.Admin.NewMessage.ApproveButton), "approve_"),
			tgbotapi.NewInlineKeyboardButtonData(m.locales.Get(locales.Messages.Admin.NewMessage.RejectButton), "reject_"),
		),
	)

	// Отправляем сообщение администратору на проверку
	msg := tgbotapi.NewMessage(m.cfg.AdminID, m.locales.Get(locales.Messages.Admin.NewMessage.MessageHeader)+userMessage)
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		return fmt.Errorf("sending message to administrator: %w", err)
	}
	return nil
}

func isMessageValid(userMessage string, minLen int) bool {
	if len(userMessage) < minLen {
		return false
	}

	return true
}
