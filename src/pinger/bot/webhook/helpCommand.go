package webhook

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewHelpCommand() Command {
	return &commandImpl{
		shortName:   "/help",
		description: "Show Help",
		onStart:     showHelp,
	}
}

func showHelp(b *BotPingers, pinger *BotPinger, message *tgbotapi.Message) {
	commandsMessage := "Available commands:\n"
	for _, c := range b.commands {
		commandsMessage += fmt.Sprintf("%s - %s\n", c.GetShortName(), c.GetDescription())
	}

	b.sendMessage(
		pinger.chatId,
		"Light Pinger Bot\n"+
			commandsMessage+
			"How it works. Bot pings provided IP every 10 minutes (configurable). If ip access status changes, bot will send you a message.",
	)
}
