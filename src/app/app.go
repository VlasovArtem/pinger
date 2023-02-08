package app

import (
	"errors"
	"github.com/spf13/pflag"
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
