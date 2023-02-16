package api

import (
	"github.com/VlasovArtem/pinger/src/pinger"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type SimplePingers struct {
	pingers map[uuid.UUID]*pinger.Pinger
}

func (SimplePingers) Init(router *mux.Router) {
	//TODO implement me
	panic("implement me")

	//subrouter := router.PathPrefix("/pinger").Subrouter()
	//
	//subrouter.Path("/add/ip").HandlerFunc(p.AddIp(false)).Methods("POST")
	//subrouter.Path("/add/ip/trusted").HandlerFunc(p.AddIp(true)).Methods("POST")
	//subrouter.Path("/consensus/{consensus}").HandlerFunc(p.SetQuorum()).Methods("PATCH")
	//subrouter.Path("/reset").HandlerFunc(p.Reset()).Methods("PATCH")
	//subrouter.Path("/timeout").HandlerFunc(p.SetTimeout()).Methods("PATCH")
	//subrouter.Path("/start").HandlerFunc(p.Start()).Methods("POST")
	//subrouter.Path("/stop").HandlerFunc(p.Stop()).Methods("PATCH")
}

//func (p *PingerApi) AddIp(trusted bool) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		ip := r.URL.Query().Get("ip")
//
//		if err := p.pinger.AddIp(ip, trusted); err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//		} else {
//			w.WriteHeader(http.StatusOK)
//		}
//	}
//}
//
//func (p *PingerApi) SetQuorum() http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		if err := p.pinger.SetQuorum(mux.Vars(request)["consensus"]); err != nil {
//			http.Error(writer, err.Error(), http.StatusBadRequest)
//		} else {
//			writer.WriteHeader(http.StatusOK)
//		}
//	}
//}
//
//func (p *PingerApi) Reset() http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		p.pinger.Reset()
//		if _, err := writer.Write([]byte("ping stopped")); err != nil {
//			http.Error(writer, err.Error(), http.StatusInternalServerError)
//		} else {
//			writer.WriteHeader(http.StatusOK)
//		}
//	}
//}
//
//func (p *PingerApi) SetTimeout() http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		query := request.URL.Query()
//
//		timeout := query.Get("timeout")
//		timeoutType := query.Get("type")
//		if err := p.pinger.SetTimeout(timeout, timeoutType); err != nil {
//			http.Error(writer, err.Error(), http.StatusBadRequest)
//		} else {
//			writer.WriteHeader(http.StatusOK)
//		}
//	}
//}
//
//func (p *PingerApi) Start() http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		if response, err := p.pinger.Start(); err != nil {
//			http.Error(writer, err.Error(), http.StatusForbidden)
//		} else {
//			if err := json.NewEncoder(writer).Encode(response); err != nil {
//				http.Error(writer, err.Error(), http.StatusInternalServerError)
//				p.pinger.Stop()
//			} else {
//				writer.WriteHeader(http.StatusOK)
//			}
//		}
//	}
//}
//
//func (p *PingerApi) Stop() http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		p.pinger.Stop()
//		writer.WriteHeader(http.StatusOK)
//	}
//}
