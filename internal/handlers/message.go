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
			tgbotapi.NewInlineKeyboardButtonData("ДА", "approve_"+userMessage),
			tgbotapi.NewInlineKeyboardButtonData("НЕТ", "reject_"+userMessage),
		),
	)

	// Отправляем сообщение администратору на проверку
	msg := tgbotapi.NewMessage(m.cfg.AdminID, "Сообщение на проверку:\n\n"+userMessage)
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		return fmt.Errorf("sending message to administrator: %w", err)
	}

	// Отправляем подтверждение пользователю
	if _, err := bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ваше сообщение добавлено в очередь на публикацию")); err != nil {
		return fmt.Errorf("sending confirmation message to user: %w", err)
	}

	return nil
}

// Обработка callback от inline кнопок
