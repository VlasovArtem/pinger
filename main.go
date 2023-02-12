package main

import (
	"github.com/VlasovArtem/pinger/src/app"
	"github.com/rs/zerolog/log"
)

func main() {
	if application, err := app.CreateApplication(); err != nil {
		log.Fatal().Msg(err.Error())
	} else {
		application.Run()
	}
}
