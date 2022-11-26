package telegram

import "fmt"

const (
	authorizeURL = "https://oauth.vk.com/authorize?client_id=%d&display=page&redirect_uri=%s&scope=%s&response_type=code&v=%s"
)

// формирование ссылки для прохождения Oauth
func (b *Bot) authorizeLink(redirectURL string) string {
	return fmt.Sprintf(authorizeURL, b.messages.AppID, redirectURL, b.messages.Scope, b.messages.Version)
}
