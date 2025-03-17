package main

import (
	"fmt"
	"log"
	"net/http"

	"tg-bot/internal/config"
	"tg-bot/internal/handlers"
	"tg-bot/internal/ui"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("load envs: %s", err.Error())
	}

	fmt.Printf("parsed config: %+v\n", cfg)

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	wh, _ := tgbotapi.NewWebhook(fmt.Sprintf("https://%s.containers.yandexcloud.net/%s", cfg.YCID, bot.Token))

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go func() {
		if errServer := http.ListenAndServe(fmt.Sprintf(":%s", cfg.YCPort), nil); err != nil {
			log.Fatal(errServer)
		}
	}()

	keyboard := ui.InitKeyboard()

	manager := handlers.NewManager(cfg)
	for update := range updates {
		if update.Message != nil { // Если это текстовое сообщение
			switch update.Message.Command() {
			case "start":
				if errUpdate := manager.StartCommand(bot, update.Message.Chat.ID, keyboard); errUpdate != nil {
					log.Printf("start message command: %s", err.Error())
				}

			case "help":
				if errUpdate := manager.HelpCommand(bot, update.Message); errUpdate != nil {
					log.Printf("help message command: %s", err.Error())
				}
			default:
				if errUpdate := manager.Message(bot, update.Message); errUpdate != nil {
					log.Printf("update message error: %s", errUpdate.Error())
				}
			}
		} else if update.CallbackQuery != nil { // Если это callback от inline-кнопки
			if errUpdate := manager.Callback(bot, update.CallbackQuery); errUpdate != nil {
				log.Printf("update message error: %s", errUpdate.Error())
			}
		} else {
			log.Printf("unknown request from chat with user: %s", update.FromChat().UserName)
		}
	}
}
