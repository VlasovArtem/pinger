package webhook

import (
	"fmt"
	"github.com/VlasovArtem/pinger/src/pinger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

func NewStatusCommand() Command {
	return &commandImpl{
		shortName:   "/status",
		description: "Show current bot status",
		onStart:     showStatus,
		condition:   botIsEnabled,
	}
}

func showStatus(pingers *BotPingers, pinger *BotPinger, message *tgbotapi.Message) {
	pingers.sendMessage(
		pinger.chatId,
		formatStatusMessage(pinger.CurrentStatus()),
	)
}

func formatStatusMessage(status pinger.PingerState) string {
	return fmt.Sprintf(
		"Bot status: %s\nLight status: %s\nTime has passed: %s",
		formatBotStatus(status.IsRunning),
		formatLightStatus(status.Pings),
		formatPassedTime(status.Pings),
	)
}

func formatPassedTime(pings []pinger.PingInfo) string {
	if len(pings) == 0 {
		return "Unknown"
	}

	lastPingInfo := pings[len(pings)-1]
	passedTime := time.Now().Sub(lastPingInfo.PingTime)
	return passedTime.String()
}

func formatBotStatus(running bool) string {
	if running {
		return "Running"
	} else {
		return "Stopped"
	}
}

func formatLightStatus(pings []pinger.PingInfo) string {
	if len(pings) == 0 {
		return "Unknown"
	}

	if pings[len(pings)-1].Result {
		return "On"
	} else {
		return "Off"
	}
}
