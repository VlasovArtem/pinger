package webhook

import (
	"github.com/VlasovArtem/pinger/src/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

func NewQuorumCommand() Command {
	return &commandImpl{
		shortName:   "/quorum",
		description: "Set quorum",
		condition:   botIsEnabled,
		onStart:     setQuorumStart,
		onContinue:  setQuorumContinue,
	}
}

func setQuorumContinue(pingers *BotPingers, pinger *BotPinger, message *tgbotapi.Message) {
	err := pinger.SetQuorum(message.Text)
	if err != nil {
		pingers.sendMessage(
			pinger.chatId,
			err.Error(),
		)
		return
	}
	setCurrentCommand(pingers, pinger, message)
}

func setQuorumStart(pingers *BotPingers, pinger *BotPinger, message *tgbotapi.Message) {
	_, err := pingers.botApi.Send(
		tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID: pinger.chatId,
				ReplyMarkup: tgbotapi.NewInlineKeyboardMarkup(
					[]tgbotapi.InlineKeyboardButton{
						tgbotapi.NewInlineKeyboardButtonData("All ips", string(config.ALL)),
						tgbotapi.NewInlineKeyboardButtonData("Any ips", string(config.ANY)),
					},
				),
			},
			Text: "Set Quorum",
		})

	if err != nil {
		log.Error().Err(err).Msg("error sending message")
	}
	unsetCurrentCommand(pingers, pinger, message)
}
