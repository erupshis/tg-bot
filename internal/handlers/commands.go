package handlers

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (m *Manager) StartCommand(bot *tgbotapi.BotAPI, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, greetingsText)
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

const greetingsText = `Привет-привет!

Я — анонимный бот поддержки для женщин в период до и после менопаузы/климакса.

❗️ Зачем меня создали?
Чтобы дать любой женщине возможность анонимно поделиться тем, что у неё наболело, и получить безоценочную поддержку.

❓ Что я делаю?
Анонимно пересылаю твоё сообщение в канал «О чем молчат женщины»:
https://t.me/menopauseanonymous

❓ Что тебе нужно сделать?
 1. Напиши в диалоге с ботом о том, что наболело. Без фильтров, без «что обо мне подумают» 🙂
 2. Нажми кнопку «Отправить» — и твоё сообщение автоматически уйдёт в канал.

❓ Как можно писать:
В любой форме. Разрешены мат, крик, вопли, КАПСЛОКИ и все, что ты чувствуешь!
Все эмоции принимаются: гнев, ярость, ненависть, скука, безразличие, апатия, грусть, раздражение (и, конечно, положительные эмоции тоже).

❗️ Одно сообщение — твоя завершённая история
Пожалуйста, выкладывай свою историю в одном сообщении, чтобы избежать путаницы.
Для удобства можешь сначала написать текст в заметках, а потом вставить сюда.

❗️ Оставляй свои контакты (ник или номер в Телеграм), если:
Тебе нужна конкретная помощь (например, при темах с суицидом, депрессией и другими сложными состояниями).
Без твоих контактов мы не сможем с тобой связаться.

❓ Не поняла/Возникла трудность/Есть вопрос или предложение:
Пиши в аккаунт технической поддержки: @kristina_rupshayte

Я постараюсь ответить как можно быстрее, но сама понимаешь, иногда нужно немного подождать 🙂

🌱 Автор проекта
Кристина Рупшайте
https://www.instagram.com/kristina_rupshayte_psy

Твои сообщения появятся на канале:
https://t.me/menopauseanonymous
👇🏼 Давай попробуем, пиши свое сообщение`
