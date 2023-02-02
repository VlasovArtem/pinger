package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func NewStartCommand() Command {
	return &commandImpl{
		shortName:   "/start",
		description: "Start pinging",
		onStart:     startPinger,
		condition:   botIsEnabled,
	}
}

func startPinger(b *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) {
	_, err := botPinger.Start()
	if err != nil {
		b.sendMessage(botPinger.chatId, err.Error())
	} else {
		b.sendMessage(botPinger.chatId, "Pinger started")
	}
}
