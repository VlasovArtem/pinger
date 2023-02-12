package api

import (
	"github.com/VlasovArtem/pinger/src/handler"
	"github.com/VlasovArtem/pinger/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const AuthHeaderName = "X-Bot-Static-Auth-Token"

func InitBotStaticApi(service service.BotStaticService, router *gin.Engine) {
	subrouter := router.Group("/bot/chat", authorized(service))

	{
		subrouter.POST("/add", addChat(service))
		subrouter.PATCH("/:id/start", startChat(service))
		subrouter.PATCH("/:id/stop", stopChat(service))
		subrouter.GET("", getChats(service))
		subrouter.GET("/:id", getChatDetails(service))
		subrouter.DELETE("/:id", deleteChat(service))
		subrouter.PUT("/:id", updateChat(service))
	}
}

func authorized(service service.BotStaticService) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader(AuthHeaderName)

		if token == "" {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		errorResponse := service.ValidateToken(token)
		if errorResponse != nil {
			context.AbortWithStatusJSON(errorResponse.Code, errorResponse)
			return
		}
	}
}

func addChat(service service.BotStaticService) gin.HandlerFunc {
	return func(context *gin.Context) {
		chatRequest := handler.AddChatRequest{}
		err := context.ShouldBind(&chatRequest)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, handler.NewBadRequestErrorResponse(err.Error()))
			return
		}
		chat, errorResponse := service.AddChat(chatRequest)
		if errorResponse != nil {
			context.AbortWithStatusJSON(errorResponse.Code, errorResponse)
			return
		}
		context.JSON(http.StatusOK, chat)
	}
}

func startChat(staticService service.BotStaticService) gin.HandlerFunc {
	return func(context *gin.Context) {
		chatIdString := context.Param("id")
		if chatIdString == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, handler.NewBadRequestErrorResponse("chat id is not set"))
			return
		} else {
			chatId, err := strconv.Atoi(chatIdString)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, handler.NewBadRequestErrorResponse(err.Error()))
				return
			}
			errorResponse := staticService.StartChat(int64(chatId))
			if errorResponse != nil {
				context.AbortWithStatusJSON(errorResponse.Code, errorResponse)
				return
			}
		}
	}
}

func stopChat(staticService service.BotStaticService) gin.HandlerFunc {
	return func(context *gin.Context) {
		chatIdString := context.Param("id")
		if chatIdString == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, handler.NewBadRequestErrorResponse("chat id is not set"))
			return
		} else {
			chatId, err := strconv.Atoi(chatIdString)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, handler.NewBadRequestErrorResponse(err.Error()))
				return
			}
			errorResponse := staticService.StopChat(int64(chatId))
			if errorResponse != nil {
				context.AbortWithStatusJSON(errorResponse.Code, errorResponse)
				return
			}
		}
	}
}

func getChats(staticService service.BotStaticService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, staticService.GetAllChats())
	}
}

func getChatDetails(staticService service.BotStaticService) gin.HandlerFunc {
	return func(context *gin.Context) {
		chatIdString := context.Param("id")
		if chatIdString == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, handler.NewBadRequestErrorResponse("chat id is not set"))
			return
		} else {
			chatId, err := strconv.Atoi(chatIdString)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, handler.NewBadRequestErrorResponse(err.Error()))
				return
			}
			details, errorResponse := staticService.GetChatDetails(int64(chatId))
			if errorResponse != nil {
				context.AbortWithStatusJSON(errorResponse.Code, errorResponse)
				return
			}
			context.JSON(http.StatusOK, details)
		}
	}
}

func deleteChat(staticService service.BotStaticService) gin.HandlerFunc {
	return func(context *gin.Context) {
		chatIdString := context.Param("id")
		if chatIdString == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, handler.NewBadRequestErrorResponse("chat id is not set"))
			return
		} else {
			chatId, err := strconv.Atoi(chatIdString)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, handler.NewBadRequestErrorResponse(err.Error()))
				return
			}
			errorResponse := staticService.DeleteChat(int64(chatId))
			if errorResponse != nil {
				context.AbortWithStatusJSON(errorResponse.Code, errorResponse)
				return
			}
			context.Status(http.StatusNoContent)
		}
	}
}

func updateChat(staticService service.BotStaticService) gin.HandlerFunc {
	return func(context *gin.Context) {
		chatIdString := context.Param("id")
		if chatIdString == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, handler.NewBadRequestErrorResponse("chat id is not set"))
			return
		} else {
			chatId, err := strconv.Atoi(chatIdString)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, handler.NewBadRequestErrorResponse(err.Error()))
				return
			}
			chatRequest := handler.UpdateChatRequest{}
			err = context.ShouldBind(&chatRequest)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, handler.NewBadRequestErrorResponse(err.Error()))
				return
			}
			response, errorResponse := staticService.UpdateChat(int64(chatId), chatRequest)
			if errorResponse != nil {
				context.AbortWithStatusJSON(errorResponse.Code, errorResponse)
				return
			}
			context.JSON(http.StatusOK, response)
		}
	}
}
