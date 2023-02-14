package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"os"
)

type Application interface {
	Run()
}

func CreateApplication() (Application, error) {
	opts := readOpts()

	switch opts.ApplicationType {
	case "bot.static":
		return NewBotStaticApplication(opts.BotStatic)
	}
	return nil, errors.New("no matching application found for type " + opts.ApplicationType)
}

func readOpts() ApplicationOpts {
	opts := ApplicationOpts{}
	pflag.StringVar(&opts.BotStatic.File, "bot.static.file", "", "File with bot static config")
	pflag.StringVar(&opts.ApplicationType, "app.type", "", "Application type")
	pflag.Parse()
	return opts
}

type ApplicationOpts struct {
	ApplicationType string
	BotStatic       BotStaticOpts
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
