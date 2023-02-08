package webhook

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func NewIpCommand() Command {
	return &commandImpl{
		shortName:   "/ip",
		description: "Set ips",
		onStart:     setIPStart,
		onContinue:  setIPContinue,
		condition:   botIsEnabled,
	}
}

func setIPStart(pingers *BotPingers, pinger *BotPinger, message *tgbotapi.Message) {
	pingers.sendMessage(
		pinger.chatId,
		"Enter ips (split by new line or comma)",
	)
	setCurrentCommand(pingers, pinger, message)
}

func setIPContinue(pingers *BotPingers, pinger *BotPinger, message *tgbotapi.Message) {
	text := message.Text
	if strings.Contains(text, ",") {
		text = strings.ReplaceAll(text, ",", "\n")
	}
	var resultMessage string
	for _, s := range strings.Split(text, "\n") {
		if s != "" {
			err := pinger.AddIp(strings.TrimSpace(s), true)
			if err != nil {
				pingers.sendMessage(
					pinger.chatId,
					err.Error(),
				)
			} else {
				resultMessage += s
			}
		}
	}
	if resultMessage != "" {
		pingers.sendMessage(
			pinger.chatId,
			"Added next ips: "+resultMessage,
		)
	}
	unsetCurrentCommand(pingers, pinger, message)
}
