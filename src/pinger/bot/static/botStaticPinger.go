package static

import (
	"github.com/VlasovArtem/pinger/src/bot"
	"github.com/VlasovArtem/pinger/src/config"
	"github.com/VlasovArtem/pinger/src/pinger"
)

type BotStaticPinger struct {
	*pinger.Pinger
}

func NewBotStaticPinger(configuration *config.PingerConfig, botApi *bot.ChatTelegramBot) *BotStaticPinger {
	return &BotStaticPinger{
		Pinger: pinger.NewPingerWithConfig(
			configuration,
			pinger.NewBotProcessor(&pinger.LightBotFormatter{}, botApi),
			pinger.NewDefaultPingProvider(),
		),
	}
}
