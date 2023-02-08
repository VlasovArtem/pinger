package app

import (
	"fmt"
	"github.com/VlasovArtem/pinger/src/bot"
	"github.com/VlasovArtem/pinger/src/config"
	"github.com/VlasovArtem/pinger/src/pinger/bot/static"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"runtime"
	"strings"
)

type botStaticApplication struct {
	config *config.BotStaticConfig
	pinger *static.BotStaticPinger
}

func NewBotStaticApplication(opts BotStaticOpts) (Application, error) {
	staticConfig, err := readBotStaticConfig(opts)
	if err != nil {
		return nil, err
	}
	pingerConfig, err := readPingerConfiguration(staticConfig)
	if err != nil {
		return nil, err
	}

	chatTelegramBot := bot.NewChatTelegramBot(
		bot.NewTelegramBot(staticConfig.Bot),
		staticConfig.Chat,
	)
	botPuller := &botStaticApplication{
		config: staticConfig,
		pinger: static.NewBotStaticPinger(
			pingerConfig,
			chatTelegramBot,
		),
	}
	sendWelcomeMessage(chatTelegramBot, pingerConfig)
	return botPuller, nil
}

func sendWelcomeMessage(telegramBot *bot.ChatTelegramBot, pingerConfig *config.PingerConfig) {
	message := fmt.Sprintf("Light Buzzer Started.\nWe will inform when %s ips [%s] will be unreachable.\nWe will check reachability within the interval %s",
		pingerConfig.Consensus,
		strings.Join(pingerConfig.Ips, ", "),
		pingerConfig.Timeout.String(),
	)
	_, err := telegramBot.SendMessage(message)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while sending message")
	}
}

func (b *botStaticApplication) Run() {
	_, err := b.pinger.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("Bot static pinger start error")
	}

	runtime.Goexit()
}

func readBotStaticConfig(opts BotStaticOpts) (*config.BotStaticConfig, error) {
	if opts.File == "" {
		return nil, errors.New("bot static config file is not set")
	}
	pullerConfig, err := config.NewBotStaticConfigFromFile(opts.File)
	if err != nil {
		return nil, errors.Wrap(err, "bot static config file read error")
	}
	return pullerConfig, nil
}

func readPingerConfiguration(botPullerConfig *config.BotStaticConfig) (*config.PingerConfig, error) {
	pingerConfig, err := config.NewPingerConfigFromFile(botPullerConfig.PingerFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "pinger config file read error")
	}
	return pingerConfig, nil
}
