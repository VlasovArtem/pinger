package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Application interface {
	Run()
	Destroy() error
}

func CreateApplication() (Application, error) {
	opts, err := readOpts()
	if err != nil {
		return nil, err
	}

	switch opts.ApplicationType {
	case "bot.static":
		return NewBotStaticApplication(opts)
	}
	return nil, errors.New("no matching application found for type " + opts.ApplicationType)
}

func readOpts() (ApplicationOpts, error) {
	opts := ApplicationOpts{}

	_, err := flags.Parse(&opts)
	if err != nil {
		return ApplicationOpts{}, err
	}
	return opts, nil
}

type ApplicationOpts struct {
	ApplicationType string `short:"t" long:"type" description:"Application type" env:"TYPE" required:"true" choice:"bot.static"`

	Api struct {
		Port string `long:"port" description:"Port for api" default:"3030" env:"PORT"`
	} `group:"api" namespace:"api" env-namespace:"API"`

	Logger struct {
		File  string `long:"file" description:"File for logger" env:"FILE"`
		Level string `long:"level" description:"Level for logger" env:"LEVEL" choice:"debug" choice:"info" choice:"warn" choice:"error" choice:"fatal" default:"info"`
	} `group:"logger" namespace:"logger" env-namespace:"LOGGER"`

	BotStatic struct {
		File string `long:"file" description:"File with bot static config" env:"FILE"`
	} `group:"bot.static" namespace:"bot.static" env-namespace:"BOT_STATIC"`
}

func PrintApi(router *gin.Engine) {
	for i, info := range router.Routes() {
		log.Info().Msgf("%d) %s %s", i, info.Method, info.Path)
	}
}

func StartRouter(router *gin.Engine, opts ApplicationOpts) {
	port := opts.Api.Port
	if port == "" {
		port = "3030"
	}

	log.Fatal().
		Err(router.Run(":" + port)).
		Msg("HTTP Application error")
}

func InitLogger(opts ApplicationOpts) (*os.File, error) {
	if opts.Logger.File != "" {
		fileLogger := &lumberjack.Logger{
			Filename:   opts.Logger.File,
			MaxSize:    1,    // megabytes
			MaxBackups: 5,    // files
			MaxAge:     7,    // days
			Compress:   true, // disabled by default
		}
		log.Debug().Msg("Logger file: " + opts.Logger.File)
		log.Logger = zerolog.New(fileLogger).With().Timestamp().Logger()
	} else if opts.Logger.Level != "" {
		level, err := zerolog.ParseLevel(opts.Logger.Level)
		if err != nil {
			return nil, err
		}
		log.Debug().Msg("Logger level: " + opts.Logger.Level)
		zerolog.SetGlobalLevel(level)
	}
	return nil, nil
}
