package main

import (
	"github.com/VlasovArtem/pinger/src/app"
	"github.com/rs/zerolog/log"
)

func main() {
	if application, err := app.CreateApplication(); err != nil {
		log.Fatal().Msg(err.Error())
	} else {
		defer func() {
			log.Info().Msg("Destroying application")
			err = application.Destroy()
			if err != nil {
				log.Fatal().Msg(err.Error())
			}
		}()

		application.Run()
	}
}
