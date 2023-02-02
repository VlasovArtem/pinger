package pinger

type ResultFormatter interface {
	FormatSuccess(prev *PingInfo, current PingInfo) string
	FormatError(prev *PingInfo, current PingInfo) string
}

type LightBotFormatter struct{}

func (LightBotFormatter) FormatSuccess(prev *PingInfo, current PingInfo) string {
	return "Light is ON. Passed time: " + formatPassedTime(prev, current)
}

func (LightBotFormatter) FormatError(prev *PingInfo, current PingInfo) string {
	return "Light is OFF. Passed time: " + formatPassedTime(prev, current)
}

func formatPassedTime(prev *PingInfo, current PingInfo) string {
	if prev == nil {
		return "0"
	}

	return current.PingTime.Sub(prev.PingTime).String()
}
