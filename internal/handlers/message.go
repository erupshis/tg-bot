package handlers

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Message Обработка текстовых сообщений
func (m *Manager) Message(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	userMessage := message.Text

	// Создаем inline клавиатуру с кнопками ДА/НЕТ
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅ Одобрить", "approve_"),
			tgbotapi.NewInlineKeyboardButtonData("❌ Отклонить", "reject_"),
		),
	)

	// Отправляем сообщение администратору на проверку
	msg := tgbotapi.NewMessage(m.cfg.AdminID, "Сообщение на проверку:\n\n"+userMessage)
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		return fmt.Errorf("sending message to administrator: %w", err)
	}

	msgChannel := tgbotapi.NewMessageToChannel(m.cfg.ChannelID, userMessage)
	msgChannel.ParseMode = "MarkdownV2"
	if _, err := bot.Send(msgChannel); err != nil {
		return fmt.Errorf("sending unchecked message in channel: %w", err)
	}

	// Отправляем подтверждение пользователю
	if _, err := bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ваше сообщение добавлено в очередь на публикацию")); err != nil {
		return fmt.Errorf("sending confirmation message to user: %w", err)
	}

	return nil
}

// Обработка callback от inline кнопок
