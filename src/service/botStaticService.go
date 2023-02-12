package service

import (
	"fmt"
	"github.com/VlasovArtem/pinger/src/bot"
	"github.com/VlasovArtem/pinger/src/config"
	"github.com/VlasovArtem/pinger/src/handler"
	"github.com/VlasovArtem/pinger/src/pinger/bot/static"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"strings"
)

type BotStaticService interface {
	AddChat(request handler.AddChatRequest) (*handler.AddChatResponse, *handler.ErrorResponse)
	StartChat(chatId int64) *handler.ErrorResponse
	StopChat(chatId int64) *handler.ErrorResponse
	GetAllChats() []handler.GetChatDetailsResponse
	GetChatDetails(chatId int64) (*handler.GetChatDetailsResponse, *handler.ErrorResponse)
	DeleteChat(chatId int64) *handler.ErrorResponse
	UpdateChat(chatId int64, request handler.UpdateChatRequest) (*handler.UpdateChatResponse, *handler.ErrorResponse)
	ValidateToken(token string) *handler.ErrorResponse
}

func NewBotStaticService(staticConfig *config.BotStaticConfig) BotStaticService {
	return &botStaticService{
		pingers:     map[int64]*static.BotStaticPinger{},
		telegramBot: bot.NewTelegramBot(staticConfig.Bot),
		config:      staticConfig,
	}
}

type botStaticService struct {
	pingers     map[int64]*static.BotStaticPinger
	telegramBot *bot.TelegramBot
	config      *config.BotStaticConfig
}

func (b *botStaticService) AddChat(request handler.AddChatRequest) (*handler.AddChatResponse, *handler.ErrorResponse) {
	log.Debug().Msgf("AddChat request: %+v", request)

	errorResponse := b.validateAddChatRequest(request)
	if errorResponse != nil {
		return nil, errorResponse
	}

	pingerConfig, err := convertConfigRequestToConfig(*request.Config)
	if err != nil {
		return nil, handler.NewBadRequestErrorResponse(err.Error())
	}

	chatTelegramBot := bot.NewChatTelegramBot(
		b.telegramBot,
		config.Chat{
			ChatId: request.ChatId,
		},
	)
	pinger := static.NewBotStaticPinger(pingerConfig,
		chatTelegramBot)
	log.Debug().Msgf("Pinger for the chat has been created [chatId: %d]", request.ChatId)
	b.pingers[request.ChatId] = pinger

	if request.AutomaticStart {
		_, err := pinger.Start()
		if err != nil {
			return nil, handler.NewForbiddenErrorResponse(err.Error())
		}
		log.Debug().Msgf("Pinger for the chat has been started [chatId: %d]", request.ChatId)
		return &handler.AddChatResponse{
			Status:  "OK",
			Details: fmt.Sprintf("Pinger for the chat has been created and started [chatId: %d]", request.ChatId),
		}, nil
	}
	err = sendWelcomeMessage(chatTelegramBot, pingerConfig)
	if err != nil {
		return nil, handler.NewForbiddenErrorResponse(err.Error())
	}
	return &handler.AddChatResponse{
		Status:  "OK",
		Details: fmt.Sprintf("Pinger for the chat has been created, but not started [chatId: %d]", request.ChatId),
	}, nil
}

func (b *botStaticService) validateAddChatRequest(request handler.AddChatRequest) *handler.ErrorResponse {
	if len(b.pingers) >= b.config.MaxNumberOfChats {
		return handler.NewForbiddenErrorResponse(fmt.Sprintf("Max number of chats [%d] exceeded", b.config.MaxNumberOfChats))
	}
	if b.pingers[request.ChatId] != nil {
		return handler.NewForbiddenErrorResponse("Chat already exists")
	}
	err := b.validateChatId(request.ChatId)
	if err != nil {
		return handler.NewBadRequestErrorResponse(err.Error())
	}
	return nil
}

func (b *botStaticService) StartChat(chatId int64) *handler.ErrorResponse {
	pinger := b.pingers[chatId]
	if pinger == nil {
		return handler.NewNotFoundErrorResponse("Pinger not found")
	}
	_, err := pinger.Start()
	if err != nil {
		return handler.NewForbiddenErrorResponse(err.Error())
	}
	return nil
}

func (b *botStaticService) StopChat(chatId int64) *handler.ErrorResponse {
	pinger := b.pingers[chatId]
	if pinger == nil {
		return handler.NewNotFoundErrorResponse("Pinger not found")
	}
	pinger.Stop()
	return nil
}

func (b *botStaticService) GetAllChats() []handler.GetChatDetailsResponse {
	var chats []handler.GetChatDetailsResponse
	for chatId, pinger := range b.pingers {
		chats = append(chats, toChatDetails(chatId, pinger))
	}
	if chats == nil {
		return []handler.GetChatDetailsResponse{}
	} else {
		return chats
	}
}

func (b *botStaticService) GetChatDetails(chatId int64) (*handler.GetChatDetailsResponse, *handler.ErrorResponse) {
	pinger := b.pingers[chatId]
	if pinger == nil {
		return nil, handler.NewNotFoundErrorResponse("Pinger not found")
	}
	details := toChatDetails(chatId, pinger)
	return &details, nil
}

func (b *botStaticService) DeleteChat(chatId int64) *handler.ErrorResponse {
	pinger := b.pingers[chatId]
	if pinger == nil {
		return handler.NewNotFoundErrorResponse("Pinger not found")
	}
	errorResponse := b.StopChat(chatId)
	if errorResponse != nil {
		return errorResponse
	}
	delete(b.pingers, chatId)
	return nil
}

func (b *botStaticService) UpdateChat(chatId int64, request handler.UpdateChatRequest) (*handler.UpdateChatResponse, *handler.ErrorResponse) {
	errorResponse := b.StopChat(chatId)
	if errorResponse != nil {
		return nil, errorResponse
	}

	response, errorResponse := b.AddChat(handler.AddChatRequest{
		ChatId:         chatId,
		Config:         request.Config,
		AutomaticStart: request.AutomaticStart,
	})
	if errorResponse != nil {
		return nil, errorResponse
	}
	return &handler.UpdateChatResponse{
		Status:  response.Status,
		Details: response.Details,
	}, nil
}

func (b *botStaticService) ValidateToken(token string) *handler.ErrorResponse {
	if token != b.config.Token {
		return handler.NewForbiddenErrorResponse("Invalid token")
	}
	return nil
}

func (b *botStaticService) validateChatId(id int64) error {
	log.Debug().Msgf("Validating chat id [%d]", id)
	chatInfo, err := b.telegramBot.GetChat(
		tgbotapi.ChatInfoConfig{
			ChatConfig: tgbotapi.ChatConfig{
				ChatID: id,
			},
		})

	if err != nil {
		return err
	}
	if chatInfo.Permissions.CanSendMessages == false {
		return errors.New("Bot has no permissions to send messages to the chat")
	}
	return nil
}

func convertConfigRequestToConfig(configRequest handler.PingerConfigRequest) (*config.PingerConfig, error) {
	log.Debug().Msgf("Converting config request to config [%v]", configRequest)
	err := validate(configRequest)
	if err != nil {
		return nil, err
	}

	return toPingerConfig(format(configRequest)), nil
}

func toChatDetails(chatId int64, pinger *static.BotStaticPinger) handler.GetChatDetailsResponse {
	currentConfig := pinger.GetCurrentConfig()

	return handler.GetChatDetailsResponse{
		ChatId: chatId,
		Config: handler.PingerConfigResponse{
			Ips:       currentConfig.Ips,
			Consensus: string(currentConfig.Consensus),
			Timeout:   currentConfig.Timeout.String(),
		},
		State: pinger.CurrentStatus(),
	}
}

func toPingerConfig(configRequest handler.PingerConfigRequest) *config.PingerConfig {
	pingerConfig := config.PingerConfig{
		Ips:       configRequest.Ips,
		Consensus: config.Consensus(configRequest.Consensus),
	}
	pingerConfig.SetTimeout(configRequest.Timeout.Value, config.TimeoutType(configRequest.Timeout.Type))
	log.Debug().Msgf("Converted config request to config [%v]", pingerConfig)
	return &pingerConfig
}

func format(configRequest handler.PingerConfigRequest) handler.PingerConfigRequest {
	configRequest.Consensus = strings.ToLower(configRequest.Consensus)
	configRequest.Timeout.Type = strings.ToLower(configRequest.Timeout.Type)
	log.Debug().Msgf("Formatted config request [%v]", configRequest)
	return configRequest
}

func validate(configRequest handler.PingerConfigRequest) error {
	if configRequest.Timeout.Value <= 0 {
		return errors.New("Timeout value must be positive")
	}
	consensus := config.Consensus(configRequest.Consensus)
	if consensus != config.ALL && consensus != config.ANY {
		return errors.New("Consensus must be 'all' or 'any'")
	}
	timeoutType := config.TimeoutType(configRequest.Timeout.Type)
	if timeoutType != config.SECONDS && timeoutType != config.MINUTES {
		return errors.New("Timeout type must be 'seconds' or 'minutes'")
	}
	if len(configRequest.Ips) == 0 {
		return errors.New("Ips can't be empty")
	}
	return nil
}

func sendWelcomeMessage(telegramBot *bot.ChatTelegramBot, pingerConfig *config.PingerConfig) error {
	log.Debug().Msg("Sending welcome message")
	message := fmt.Sprintf("Light Buzzer Started.\nWe will inform when %s ips [%s] will be unreachable.\nWe will check reachability within the interval %s",
		pingerConfig.Consensus,
		strings.Join(pingerConfig.Ips, ", "),
		pingerConfig.Timeout.String(),
	)
	_, err := telegramBot.SendMessage(message)
	if err != nil {
		return errors.New("Error while sending welcome message")
	}
	return nil
}
