package telegram

import (
	"fmt"

	"github.com/AlexKomzzz/tlg-bot-VK/pkg/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// авторизация клиента
func (b *Bot) initAuthorizationProcess(message *tgbotapi.Message) error {
	// формирование токена авторизации от VK API
	authLink := b.createAuthorizationLink(message.Chat.ID)

	// отправка клиенту ссылку для прохождения авторизации
	msgText := fmt.Sprintf(b.messages.Responses.Start, authLink)
	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) createAuthorizationLink(chatID int64) string {
	// формирование ссылки редиректа
	redirectUrl := b.generateRedirectURL(chatID)

	return b.authorizeLink(redirectUrl)
}

// добавление в ссылку редиректа chatID
func (b *Bot) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.messages.RedirectURL, chatID)
}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.storage.Get(chatID, storage.AccessTokens)
}
