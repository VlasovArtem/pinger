package bot

import (
	"github.com/VlasovArtem/pinger/src/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ChatTelegramBot struct {
	*TelegramBot
	reference config.Chat
}

func NewChatTelegramBot(bot *TelegramBot, reference config.Chat) *ChatTelegramBot {
	return &ChatTelegramBot{TelegramBot: bot, reference: reference}
}

func (c *ChatTelegramBot) SendMessage(msg string) (tgbotapi.Message, error) {
	var message tgbotapi.MessageConfig
	if c.reference.Username == "" {
		message = tgbotapi.NewMessage(c.reference.ChatId, msg)
	} else {
		message = tgbotapi.NewMessageToChannel(c.reference.Username, msg)
	}
	return c.Send(message)
}
