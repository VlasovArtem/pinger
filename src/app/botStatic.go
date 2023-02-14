package app

import (
	"github.com/VlasovArtem/pinger/src/api"
	"github.com/VlasovArtem/pinger/src/config"
	"github.com/VlasovArtem/pinger/src/service"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type botStaticApplication struct {
	service service.BotStaticService
	opts    ApplicationOpts
	destroy func() error
}

func NewBotStaticApplication(opts ApplicationOpts) (Application, error) {
	file, err := InitLogger(opts)
	if err != nil {
		return nil, err
	}
	staticConfig, err := readBotStaticConfig(opts)
	if err != nil {
		return nil, err
	}

	botPuller := &botStaticApplication{
		service: service.NewBotStaticService(staticConfig),
		opts:    opts,
		destroy: func() error {
			return file.Close()
		},
	}
	return botPuller, nil
}

func (b *botStaticApplication) Run() {
	engine := gin.Default()
	err := engine.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal().Err(err).Msg("SetTrustedProxies error")
	}

	api.InitBotStaticApi(b.service, engine)

	PrintApi(engine)

	StartRouter(engine, b.opts)
}

func (b *botStaticApplication) Destroy() error {
	return b.destroy()
}

func readBotStaticConfig(opts ApplicationOpts) (*config.BotStaticConfig, error) {
	if opts.BotStatic.File == "" {
		return nil, errors.New("bot static config file is not set")
	}
	pullerConfig, err := config.NewBotStaticConfigFromFile(opts.BotStatic.File)
	if err != nil {
		return nil, errors.Wrap(err, "bot static config file read error")
	}
	return pullerConfig, nil
}
