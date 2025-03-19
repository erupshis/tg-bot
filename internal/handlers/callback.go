package handlers

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (m *Manager) Callback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) error {
	// Разделяем callback_data на действие и сообщение
	data := strings.SplitN(callback.Data, "_", 2)
	action := data[0]
	//userMessage := data[1]

	if action == "approve" {
		// Отправляем сообщение в канал
		//msg := tgbotapi.NewMessageToChannel(m.cfg.ChannelID, userMessage)
		//msg.ParseMode = "MarkdownV2"
		//if _, err := bot.Send(msg); err != nil {
		//	return fmt.Errorf("sending approved message in channel: %w", err)
		//}

		// Редактируем сообщение администратора
		edit := tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, "Сообщение одобрено и отправлено в канал.")
		if _, err := bot.Send(edit); err != nil {
			return fmt.Errorf("modify approved message at admin's cchat: %w", err)
		}
	} else if action == "reject" {
		// Редактируем сообщение администратора
		edit := tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, "Сообщение отклонено и не будет отправлено в канал.")
		if _, err := bot.Send(edit); err != nil {
			return fmt.Errorf("modify rejected message at admin's cchat: %w", err)
		}
	}

	return nil
}
