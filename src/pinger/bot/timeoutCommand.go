package bot

import (
	"fmt"
	"github.com/VlasovArtem/pinger/src/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewTimeoutCommands() []Command {
	return []Command{
		createTimeout(
			"/timeout", "Set timeout",
			setTimeoutStart,
			setTimeoutContinue,
		),
		createTimeout(
			"/timeout10m", "Set timeout 10 minutes",
			setTimeout("10", config.MINUTES),
			nil,
		),
		createTimeout(
			"/timeout5m", "Set timeout 5 minutes",
			setTimeout("5", config.MINUTES),
			nil,
		),
		createTimeout(
			"/timeout1m", "Set timeout 1 minutes",
			setTimeout("1", config.MINUTES),
			nil,
		),
	}
}

func createTimeout(shortName string, description string, onStart commandFunc, onContinue commandFunc) *commandImpl {
	return &commandImpl{
		shortName:   shortName,
		description: description,
		condition:   botIsEnabled,
		onStart:     onStart,
		onContinue:  onContinue,
	}
}

func setTimeoutStart(pingers *BotPingers, pinger *BotPinger, message *tgbotapi.Message) {
	pingers.sendMessage(
		pinger.chatId,
		"Enter timeout in seconds",
	)
	setCurrentCommand(pingers, pinger, message)
}

func setTimeoutContinue(pingers *BotPingers, pinger *BotPinger, message *tgbotapi.Message) {
	err := pinger.SetTimeout(message.Text, string(config.SECONDS))
	if err != nil {
		pingers.sendMessage(
			pinger.chatId,
			fmt.Sprintf("Error: %s", err),
		)
		return
	}
	unsetCurrentCommand(pingers, pinger, message)
}

func setTimeout(timeout string, timeoutType config.TimeoutType) commandFunc {
	return func(b *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) {
		botPinger.SetTimeout(timeout, string(timeoutType))
		b.sendMessage(botPinger.chatId, fmt.Sprintf("Timeout set to %s %s", timeout, timeoutType))
	}
}
