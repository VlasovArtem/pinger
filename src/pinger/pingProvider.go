package pinger

import (
	"errors"
	ping "github.com/prometheus-community/pro-bing"
	"github.com/rs/zerolog/log"
	"time"
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
	log.Debug().Msgf("Pinging %s", address)
	pinger, err := ping.NewPinger(address)
	if err != nil {
		return err
	}
	pinger.Count = 5
	pinger.Timeout = pinger.Interval * (time.Duration)(pinger.Count)
	log.Debug().Msg("Pinger started")
	err = pinger.Run()
	if err != nil {
		return err
	}
	log.Debug().Msg("Pinger stopped")
	if statistics := pinger.Statistics(); statistics == nil {
		return errors.New("statistics is nil")
	} else {
		if statistics.PacketsRecv == 0 {
			return errors.New("no packets received")
		}
	}
	return nil
}
