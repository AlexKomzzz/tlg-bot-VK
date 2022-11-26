package telegram

import (
	"github.com/AlexKomzzz/tlg-bot-VK/pkg/config"
	"github.com/AlexKomzzz/tlg-bot-VK/pkg/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI

	storage storage.TokenStorage

	messages config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, storage storage.TokenStorage, messages config.Messages) *Bot {
	return &Bot{
		bot:      bot,
		storage:  storage,
		messages: messages,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// Handle commands
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}

			continue
		}

		// Handle regular messages
		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}

	return nil
}
