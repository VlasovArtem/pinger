package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
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
	pflag.Parse()
	return opts
}

type ApplicationOpts struct {
	ApplicationType string
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

func PrintApi(router *gin.Engine) {
	for i, info := range router.Routes() {
		log.Info().Msgf("%d) %s %s", i, info.Method, info.Path)
	}
}

func StartRouter(router *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
	}

	log.Fatal().
		Err(router.Run(":" + port)).
		Msg("HTTP Application error")
}

func InitLogger(logger Logger) (*os.File, error) {
	if logger.File != "" {
		file, err := os.OpenFile(
			"myapp.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			return nil, err
		}
		log.Logger = zerolog.New(file).With().Timestamp().Logger()
	} else if logger.Level != "" {
		level, err := zerolog.ParseLevel(logger.Level)
		if err != nil {
			return nil, err
		}
		zerolog.SetGlobalLevel(level)
	}
	return nil, nil
}
