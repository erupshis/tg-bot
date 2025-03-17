package ui

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitKeyboard() tgbotapi.ReplyKeyboardMarkup {
	// Создаем персистентную Reply-клавиатуру
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/start"),
			tgbotapi.NewKeyboardButton("/help"),
		),
	)
	keyboard.OneTimeKeyboard = false // Клавиатура будет всегда видна
	keyboard.ResizeKeyboard = true

	return keyboard
}
