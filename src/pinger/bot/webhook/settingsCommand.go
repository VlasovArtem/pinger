package webhook

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"time"
)

func NewSettingsCommand() Command {
	return &commandImpl{
		shortName:   "/settings",
		description: "Show current bot settings",
		onStart:     showSettings,
		condition:   botIsEnabled,
	}
}

func showSettings(pingers *BotPingers, pinger *BotPinger, message *tgbotapi.Message) {
	pingers.sendMessage(
		pinger.chatId,
		pinger.printSettings(),
	)
}

func (p *BotPinger) printSettings() string {
	config := p.GetCurrentConfig()
	p.CurrentState()
	return fmt.Sprintf("Current settings:\n"+
		"Timeout: %d seconds\n"+
		"Quorum: %s\n"+
		"Ips: %s\n",
		config.Timeout/time.Second,
		config.Quorum,
		strings.Join(config.Ips, ", "))
}
