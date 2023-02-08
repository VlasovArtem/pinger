package api

//type Bot struct {
//	pingers *webhook.BotPingers
//}
//
//func NewBot() *Bot {
//	return &Bot{webhook.NewBotPingers()}
//}
//
//func (b *Bot) Init(router *mux.Router) {
//	router.Path("/status").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		writer.WriteHeader(http.StatusOK)
//		writer.Write([]byte("ok"))
//	})
//	subrouter := router.PathPrefix("/webhook").Subrouter()
//
//	subrouter.Path("/updates").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		var update tgbotapi.Update
//		err := json.NewDecoder(request.Body).Decode(&update)
//		if err != nil {
//			log.Err(err).Msg("update decode error")
//		} else {
//			b.pingers.PerformUpdate(update)
//		}
//	}).Methods("POST")
//}
