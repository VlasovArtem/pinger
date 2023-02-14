package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Application interface {
	Run()
	Destroy() error
}

func CreateApplication() (Application, error) {
	opts := readOpts()

	switch opts.ApplicationType {
	case "bot.static":
		return NewBotStaticApplication(opts)
	}
	return nil, errors.New("no matching application found for type " + opts.ApplicationType)
}

func readOpts() ApplicationOpts {
	opts := ApplicationOpts{}
	pflag.StringVar(&opts.BotStatic.File, "bot.static.file", "", "File with bot static config")
	pflag.StringVar(&opts.ApplicationType, "app.type", "", "Application type")
	pflag.StringVar(&opts.Logger.File, "logger.file", "", "File for logger")
	pflag.StringVar(&opts.Logger.Level, "logger.level", "", "Level for logger")
	pflag.StringVar(&opts.Api.Port, "api.port", "", "Port for api")
	pflag.Parse()
	return opts
}

type ApplicationOpts struct {
	ApplicationType string
	Api             ApiOpts
	Logger          Logger
	BotStatic       BotStaticOpts
}
type Logger struct {
	File  string
	Level string
}

type BotStaticOpts struct {
	File string
}

type ApiOpts struct {
	Port string
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

func InitLogger(logger Logger) (*os.File, error) {
	if logger.File != "" {
		fileLogger := &lumberjack.Logger{
			Filename:   logger.File,
			MaxSize:    1,    // megabytes
			MaxBackups: 5,    // files
			MaxAge:     7,    // days
			Compress:   true, // disabled by default
		}
		log.Logger = zerolog.New(fileLogger).With().Timestamp().Logger()
	} else if logger.Level != "" {
		level, err := zerolog.ParseLevel(logger.Level)
		if err != nil {
			return nil, err
		}
		zerolog.SetGlobalLevel(level)
	}
	return nil, nil
}
