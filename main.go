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
	//router := mux.NewRouter().StrictSlash(true)
	//
	//http.Handle("/", router)
	//
	//bot := api.NewBot()
	//bot.Init(router)
	//
	//if err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	//	if template, err := route.GetPathTemplate(); err != nil {
	//		log.Error().Err(err)
	//	} else {
	//		log.Info().Msg(template)
	//	}
	//	return nil
	//}); err != nil {
	//	log.Fatal().Err(err).Msg("router walk error")
	//}
	//
	//log.Fatal().
	//	Err(http.ListenAndServe(":3030", router)).
	//	Msg("HTTP Application error")
}
