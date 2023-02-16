package pinger

import (
	"fmt"
	"time"
)

type ResultFormatter interface {
	FormatSuccess(prev *PingInfo, current PingInfo) string
	FormatError(prev *PingInfo, current PingInfo) string
}

type LightBotFormatter struct{}

func (LightBotFormatter) FormatSuccess(prev *PingInfo, current PingInfo) string {
	return fmt.Sprintf(`Light is @bold@ON@bold@.
Passed time: %s (±%s)`, formatPassedTime(prev, current), current.Config.Timeout.String())
}

func (LightBotFormatter) FormatError(prev *PingInfo, current PingInfo) string {
	return fmt.Sprintf(`Light is @bold@OFF@bold@.
Passed time: %s (±%s)`, formatPassedTime(prev, current), current.Config.Timeout.String())
}

func formatPassedTime(prev *PingInfo, current PingInfo) string {
	var previousTime time.Time
	if prev != nil {
		previousTime = prev.PingTime
	} else {
		previousTime = current.PingTime
	}

	var currentTime = current.PingTime

	return currentTime.Sub(previousTime).Truncate(time.Second).String()
}
