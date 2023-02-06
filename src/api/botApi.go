package api

import (
	"encoding/json"
	"github.com/VlasovArtem/pinger/src/pinger/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Bot struct {
	pingers *bot.BotPingers
}

func NewBot() *Bot {
	return &Bot{bot.NewBotPingers()}
}

func (b *Bot) Init(router *mux.Router) {
	router.Path("/status").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("ok"))
	})
	subrouter := router.PathPrefix("/webhook").Subrouter()

	subrouter.Path("/updates").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var update tgbotapi.Update
		err := json.NewDecoder(request.Body).Decode(&update)
		if err != nil {
			log.Err(err).Msg("update decode error")
		} else {
			b.pingers.PerformUpdate(update)
		}
	}).Methods("POST")
}
