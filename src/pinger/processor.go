package pinger

import (
	"github.com/VlasovArtem/pinger/src/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
	GetDefaultConfig() *config.Config
}

type BotProcessor struct {
	resultFormatter ResultFormatter
	botAPI          *tgbotapi.BotAPI
	chatId          int64
}

func NewBotProcessor(formatter ResultFormatter, api *tgbotapi.BotAPI, chatId int64) Processor {
	return &BotProcessor{
		resultFormatter: formatter,
		botAPI:          api,
		chatId:          chatId,
	}
}

func (p *BotProcessor) OnSuccess(prev *PingInfo, current PingInfo) {
	p.sendMessage(p.resultFormatter.FormatSuccess(prev, current))
}

func (p *BotProcessor) sendMessage(message string) {
	_, err := p.botAPI.Send(tgbotapi.NewMessage(p.chatId, message))
	if err != nil {
		log.Err(err).Msg("send message error")
	}
}

func (p *BotProcessor) OnError(prev *PingInfo, current PingInfo) {
	p.sendMessage(p.resultFormatter.FormatError(prev, current))
}

func (p *BotProcessor) GetTrigger() Trigger {
	return ON_CHANGE
}

func (p *BotProcessor) GetDefaultConfig() *config.Config {
	return config.NewConfig(config.ANY, time.Minute*10)
}
