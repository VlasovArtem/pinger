package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

func NewEnableCommand() Command {
	return &commandImpl{
		shortName:   "/enable",
		description: "Enable bot (required before start)",
		onStart:     enablePingerStart,
		onContinue:  enablePingerContinue,
	}
}

func enablePingerStart(botPingers *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) {
	if botPinger.enabled {
		botPingers.sendMessage(botPinger.chatId, "Bot is already enabled")
	} else {
		botPingers.sendMessage(botPinger.chatId, "Please input enabling token")
		setCurrentCommand(botPingers, botPinger, message)
	}
}

func enablePingerContinue(b *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) {
	if message.Text == b.enablingToken {
		botPinger.enabled = true
		_, err := b.botApi.Send(&tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID: message.Chat.ID,
			},
			Text: "Bot is enabled",
		})
		if err != nil {
			log.Err(err).Msg("send message error")
		}
		params := tgbotapi.Params{}
		params.AddNonZero64("chat_id", message.Chat.ID)
		err = params.AddInterface("menu_button", map[string]string{
			"type": "commands",
		})
		if err != nil {
			log.Err(err).Msg("add menu button error")
		} else {
			_, err := b.botApi.MakeRequest("setChatMenuButton", params)
			if err != nil {
				log.Err(err).Msg("set chat menu button error")
			}
		}
	} else {
		_, err := b.botApi.Send(&tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID: message.Chat.ID,
			},
			Text: "Enabling token is not correct",
		})
		if err != nil {
			log.Err(err).Msg("send message error")
		}
	}
	unsetCurrentCommand(b, botPinger, message)
}
