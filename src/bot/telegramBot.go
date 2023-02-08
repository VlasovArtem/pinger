package bot

import (
	"github.com/VlasovArtem/pinger/src/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type TelegramBot struct {
	*tgbotapi.BotAPI
	config config.BotConfig
}

func NewTelegramBot(botConfig config.BotConfig) *TelegramBot {
	bot := &TelegramBot{config: botConfig}
	bot.init()
	return bot
}

func (b *TelegramBot) init() {
	var api *tgbotapi.BotAPI
	var err error
	if b.config.Url == "" {
		api, err = tgbotapi.NewBotAPI(b.config.Token)
	} else {
		api, err = tgbotapi.NewBotAPIWithAPIEndpoint(b.config.Token, b.config.Url)
	}
	if err != nil {
		log.Fatal().Err(err).Msg("bot api init error")
	}
	b.BotAPI = api
}
