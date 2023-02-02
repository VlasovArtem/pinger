package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func NewResetCommand() Command {
	return &commandImpl{
		shortName:   "/reset",
		description: "Reset pinger",
		onStart:     resetPinger,
		condition:   botIsEnabled,
	}
}

func resetPinger(pingers *BotPingers, pinger *BotPinger, message *tgbotapi.Message) {
	pinger.Reset()
	pingers.sendMessage(pinger.chatId, "Pinger reset")
}
