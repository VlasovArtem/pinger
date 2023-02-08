package pinger

import (
	"errors"
	"fmt"
	"github.com/VlasovArtem/pinger/src/config"
	"github.com/rs/zerolog/log"
	"net"
	"strconv"
	"strings"
)

type PingerStatus struct {
	IsRunning bool
	Pings     []PingInfo
}

type Pinger struct {
	currentConfig *config.PingerConfig
	processor     Processor
	starter       Executor
	pingProvider  PingProvider
	pings         []PingInfo
}

func NewPinger(processor Processor, provider PingProvider) *Pinger {
	return NewPingerWithConfig(processor.GetDefaultConfig(), processor, provider)
}

func NewPingerWithConfig(currentConfig *config.PingerConfig, processor Processor, provider PingProvider) *Pinger {
	return &Pinger{
		currentConfig: currentConfig,
		processor:     processor,
		pingProvider:  provider,
		starter:       NewStarter(),
		pings:         []PingInfo{},
	}
}

func (p *Pinger) GetCurrentConfig() config.PingerConfig {
	return *p.currentConfig
}

func (p *Pinger) AddIp(ip string, trusted bool) error {
	err := validateIp(ip)
	if err != nil {
		return err
	} else {
		if trusted {
			p.currentConfig.AddIp(ip)
		} else {
			err := p.pingProvider.Ping(ip)
			if err != nil {
				errorMsg := fmt.Sprintf("Ip '%s' is not responding", ip)
				log.Err(err).Msg(errorMsg)
				return err
			} else {
				p.currentConfig.AddIp(ip)
			}
		}
	}
	return nil
}

func validateIp(ip string) error {
	parseIP := net.ParseIP(ip)
	if parseIP == nil {
		return errors.New(fmt.Sprintf("Ip '%s' is not valid", ip))
	}
	return nil
}

func (p *Pinger) SetConsensus(cons string) error {
	consensus := config.Consensus(strings.ToLower(cons))
	if config.ALL != consensus && config.ANY != consensus {
		return errors.New("consensus is not valid. Should be 'ALL' or ANY")
	}
	p.currentConfig.SetConsensus(consensus)
	return nil
}

func (p *Pinger) Reset() {
	p.currentConfig = p.processor.GetDefaultConfig()
	p.Stop()
}

func (p *Pinger) SetTimeout(timeout string, timeoutType string) error {
	if timeout, err := parseTimeout(timeout); err != nil {
		return err
	} else {
		if timeoutType, err := parseTimeoutType(timeoutType); err != nil {
			return err
		} else {
			p.currentConfig.SetTimeout(timeout, timeoutType)
			return nil
		}
	}
}

func parseTimeout(timeoutString string) (int64, error) {
	if timeout, err := strconv.Atoi(timeoutString); err != nil {
		return 0, errors.New("timeout is not integer")
	} else if timeout < 1 {
		return 0, errors.New("timeout should be greater then zero")
	} else {
		return int64(timeout), nil
	}
}

func parseTimeoutType(timeoutType string) (config.TimeoutType, error) {
	if timeoutType == "" {
		return "", errors.New("timeout type could not be empty")
	}
	lowerTimeoutType := (config.TimeoutType)(strings.ToLower(timeoutType))
	if config.MINUTES != lowerTimeoutType && config.SECONDS != lowerTimeoutType {
		return "", errors.New(fmt.Sprintf("timeout type is not exists. Valid types: '%s', '%s'", config.MINUTES, config.SECONDS))
	}
	return lowerTimeoutType, nil
}

func (p *Pinger) Start() (any, error) {
	if err := validateCurrentConfig(p.currentConfig); err != nil {
		return nil, err
	} else {
		response := struct {
			message       string
			currentConfig config.PingerConfig
		}{
			message:       "Pinger started",
			currentConfig: *p.currentConfig,
		}

		p.starter.Start(
			p.getCurrentConfig,
			p.runPing,
			p.runResult,
		)

		return response, nil
	}
}

func validateCurrentConfig(currentConfig *config.PingerConfig) error {
	if len(currentConfig.GetIps()) <= 0 {
		return errors.New("ips are not exists")
	}
	return nil
}

func (p *Pinger) getCurrentConfig() config.PingerConfig {
	return *p.currentConfig
}

func (p *Pinger) runPing(runConfig config.PingerConfig) bool {
	for _, ip := range runConfig.GetIps() {
		err := p.pingProvider.Ping(ip)
		if err != nil {
			log.Err(err).Msg(fmt.Sprintf("Ip '%s' is not responding", ip))
		}
		if err != nil && runConfig.IsAll() {
			return false
		} else if err == nil && runConfig.IsAny() {
			return true
		}
	}
	if runConfig.IsAny() {
		return false
	} else {
		return true
	}
}

func (p *Pinger) runResult(info PingInfo) {
	if len(p.pings) == 0 {
		p.runProcessor(nil, info)
	} else {
		lastPing := p.pings[len(p.pings)-1]
		if lastPing.Result != info.Result {
			p.runProcessor(&lastPing, info)
		} else if p.processor.GetTrigger() == CONSTANTLY {
			p.runProcessor(&lastPing, info)
		}
	}

	if len(p.pings) == 10 {
		p.pings = p.pings[1:]
	}
	p.pings = append(p.pings, info)
}

func (p *Pinger) Stop() {
	p.starter.Stop()
}

func (p *Pinger) runProcessor(previousResult *PingInfo, currentResult PingInfo) {
	if currentResult.Result {
		p.processor.OnSuccess(previousResult, currentResult)
	} else {
		p.processor.OnError(previousResult, currentResult)
	}
}

func (p *Pinger) CurrentStatus() PingerStatus {
	return PingerStatus{
		IsRunning: p.starter.Status(),
		Pings:     p.pings,
	}
}
