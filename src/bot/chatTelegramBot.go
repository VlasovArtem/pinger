package bot

import (
	"github.com/VlasovArtem/pinger/src/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const defaultParseMode = tgbotapi.ModeMarkdownV2

var specialMappers = map[string]string{
	"@bold@": "*",
}

type ChatTelegramBot struct {
	*TelegramBot
	reference config.Chat
}

func NewChatTelegramBot(bot *TelegramBot, reference config.Chat) *ChatTelegramBot {
	return &ChatTelegramBot{TelegramBot: bot, reference: reference}
}

func (c *ChatTelegramBot) SendMessage(msg string) (tgbotapi.Message, error) {
	var message tgbotapi.MessageConfig
	sendMessage := prepareMessage(msg)
	if c.reference.Username == "" {
		message = tgbotapi.NewMessage(c.reference.ChatId, sendMessage)
	} else {
		message = tgbotapi.NewMessageToChannel(c.reference.Username, sendMessage)
	}
	message.ParseMode = defaultParseMode
	return c.Send(message)
}

func prepareMessage(msg string) string {
	escapedMessage := tgbotapi.EscapeText(defaultParseMode, msg)

	for special, markdown := range specialMappers {
		escapedMessage = strings.ReplaceAll(escapedMessage, special, markdown)
	}

	return escapedMessage
}
