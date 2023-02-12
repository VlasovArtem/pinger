package handler

import "github.com/VlasovArtem/pinger/src/pinger"

type AddChatRequest struct {
	ChatId         int64                `json:"chat_id"`
	AutomaticStart bool                 `json:"automatic_start"`
	Config         *PingerConfigRequest `json:"config"`
}

type PingerConfigRequest struct {
	Ips       []string
	Consensus string
	Timeout   struct {
		Value int64
		Type  string
	}
}

type AddChatResponse struct {
	Status  string `json:"status"`
	Details string `json:"details"`
}

type GetChatDetailsResponse struct {
	ChatId int64                `json:"chat_id"`
	Config PingerConfigResponse `json:"config"`
	State  pinger.PingerStatus  `json:"state"`
}

type PingerConfigResponse struct {
	Ips       []string
	Consensus string
	Timeout   string
}

type UpdateChatRequest struct {
	AutomaticStart bool                 `json:"automatic_start"`
	Config         *PingerConfigRequest `json:"config"`
}

type UpdateChatResponse struct {
	Status  string `json:"status"`
	Details string `json:"details"`
}
