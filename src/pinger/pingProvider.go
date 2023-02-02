package pinger

import (
	"errors"
	ping "github.com/prometheus-community/pro-bing"
)

type PingProvider interface {
	Ping(element string) error
}

type DefaultPingProvider struct {
}

func NewDefaultPingProvider() PingProvider {
	return new(DefaultPingProvider)
}

func (d *DefaultPingProvider) Ping(address string) error {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		return err
	}
	pinger.Count = 3
	pinger.SetPrivileged(true)
	err = pinger.Run()
	if err != nil {
		return err
	}
	if statistics := pinger.Statistics(); statistics == nil {
		return errors.New("statistics is nil")
	} else {
		if statistics.PacketsRecv == 0 {
			return errors.New("no packets received")
		}
	}
	return nil
}
