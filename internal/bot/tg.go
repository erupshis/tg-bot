package bot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type UpdateManager interface {
	Callback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) error

	Message(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error

	StartCommand(bot *tgbotapi.BotAPI, chatID int64) error
	HelpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error
}

type Telegram struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramBot(token, ycID string, debug bool) (*Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("init botAPI: %w", err)
	}

	bot.Debug = debug
	webhookLink := fmt.Sprintf("https://%s.containers.yandexcloud.net/%s", ycID, token)
	wh, err := tgbotapi.NewWebhook(webhookLink)
	if err != nil {
		return nil, fmt.Errorf("init webhook: %w", err)
	}

	_, err = bot.Request(wh)
	if err != nil {
		return nil, fmt.Errorf("test request: %w", err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		return nil, fmt.Errorf("get webhook info: %w", err)
	}

	if info.LastErrorDate != 0 {
		return nil, fmt.Errorf("telegram callback failed: %s", info.LastErrorMessage)
	}

	return &Telegram{
		bot: bot,
	}, nil
}

func (tg *Telegram) Run(ctx context.Context, manager UpdateManager) error {
	updates := tg.bot.ListenForWebhook("/" + tg.bot.Token)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case update := <-updates:
			// Callback от кнопок
			if update.CallbackQuery != nil {
				if errUpdate := manager.Callback(tg.bot, update.CallbackQuery); errUpdate != nil {
					logrus.Errorf("update message error: %s", errUpdate.Error())
				}
			}

			// Текстовое сообщение
			if update.Message != nil {
				if errUpdate := tg.handleMessage(update, manager); errUpdate != nil {
					logrus.Errorf("handleMessage error: %s", errUpdate.Error())
				}
			}
		}
	}
}

const (
	commandStart = "start"
	commandHelp  = "help"
)

func (tg *Telegram) handleMessage(update tgbotapi.Update, manager UpdateManager) error {
	switch update.Message.Command() {
	case commandStart:
		if err := manager.StartCommand(tg.bot, update.Message.Chat.ID); err != nil {
			return fmt.Errorf("%s command: %w", commandStart, err)
		}

	case commandHelp:
		if err := manager.HelpCommand(tg.bot, update.Message); err != nil {
			return fmt.Errorf("%s command: %w", commandHelp, err)
		}
	default:
		if err := manager.Message(tg.bot, update.Message); err != nil {
			return fmt.Errorf("message: %w", err)
		}
	}

	return nil
}
