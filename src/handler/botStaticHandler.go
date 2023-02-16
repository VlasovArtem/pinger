package handler

import (
	"time"
)

type AddChatRequest struct {
	ChatId         int64                `json:"chat_id" binding:"required,ne=0"`
	AutomaticStart bool                 `json:"automatic_start"`
	Config         *PingerConfigRequest `json:"config" binding:"required"`
}

type PingerConfigRequest struct {
	Ips     []string `json:"ips" binding:"required,len=1"`
	Quorum  string   `json:"quorum" binding:"required,oneof=all any"`
	Timeout struct {
		Value int64  `json:"value" binding:"required,gte=1"`
		Type  string `json:"type" binding:"required,oneof=seconds minutes"`
	} `json:"timeout" binding:"required"`
}

type AddChatResponse struct {
	Status  string `json:"status"`
	Details string `json:"details"`
}

type GetChatDetailsResponse struct {
	ChatId int64                `json:"chat_id"`
	Config PingerConfigResponse `json:"config"`
	State  PingerStateResponse  `json:"state"`
}

type PingerStateResponse struct {
	IsRunning         bool
	PingHistory       []PingInfoResponse
	PingChangeHistory []PingInfoResponse
}

type PingInfoResponse struct {
	Config   PingerConfigResponse
	Result   bool
	PingTime time.Time
}

type PingerConfigResponse struct {
	Ips     []string `json:"ips"`
	Quorum  string   `json:"quorum"`
	Timeout string   `json:"timeout"`
}

type UpdateChatRequest struct {
	AutomaticStart bool                 `json:"automatic_start"`
	Config         *PingerConfigRequest `json:"config" binding:"required"`
}

type UpdateChatResponse struct {
	Status  string `json:"status"`
	Details string `json:"details"`
}
