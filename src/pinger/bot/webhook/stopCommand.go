package webhook

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func NewStopCommand() Command {
	return &commandImpl{
		shortName:   "/stop",
		description: "Stop pinging",
		onStart:     stopPinger,
		condition:   botIsEnabled,
	}
}

func stopPinger(b *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) {
	botPinger.Stop()
	b.sendMessage(botPinger.chatId, "Pinger stopped")
}
