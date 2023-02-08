package pinger

import (
	"github.com/VlasovArtem/pinger/src/bot"
	"github.com/VlasovArtem/pinger/src/config"
	"github.com/rs/zerolog/log"
	"time"
)

type Trigger string

const (
	ON_CHANGE  Trigger = "on_change"
	CONSTANTLY Trigger = "constantly"
)

type Processor interface {
	OnSuccess(prev *PingInfo, current PingInfo)
	OnError(prev *PingInfo, current PingInfo)
	GetTrigger() Trigger
	GetDefaultConfig() *config.PingerConfig
}

type BotProcessor struct {
	resultFormatter ResultFormatter
	chatTelegramBot *bot.ChatTelegramBot
}

func NewBotProcessor(formatter ResultFormatter, chatTelegramBot *bot.ChatTelegramBot) Processor {
	return &BotProcessor{
		resultFormatter: formatter,
		chatTelegramBot: chatTelegramBot,
	}
}

func (p *BotProcessor) OnSuccess(prev *PingInfo, current PingInfo) {
	_, err := p.chatTelegramBot.SendMessage(p.resultFormatter.FormatSuccess(prev, current))
	if err != nil {
		log.Fatal().Err(err).Msg("Error while sending message")
	}
}

func (p *BotProcessor) OnError(prev *PingInfo, current PingInfo) {
	_, err := p.chatTelegramBot.SendMessage(p.resultFormatter.FormatError(prev, current))
	if err != nil {
		log.Fatal().Err(err).Msg("Error while sending message")
	}
}

func (p *BotProcessor) GetTrigger() Trigger {
	return ON_CHANGE
}

func (p *BotProcessor) GetDefaultConfig() *config.PingerConfig {
	return config.NewConfig(config.ANY, time.Minute*10)
}
