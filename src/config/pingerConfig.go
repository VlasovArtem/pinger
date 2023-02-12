package config

import (
	"time"
)

type Consensus string

const (
	ALL Consensus = "all"
	ANY Consensus = "any"
)

type TimeoutType string

const (
	SECONDS TimeoutType = "seconds"
	MINUTES TimeoutType = "minutes"
)

type PingerConfig struct {
	Ips       []string
	Consensus Consensus
	Timeout   time.Duration
}

func NewConfig(consensus Consensus, timeout time.Duration) *PingerConfig {
	return &PingerConfig{
		Ips:       []string{},
		Consensus: consensus,
		Timeout:   timeout,
	}
}

func (c *PingerConfig) AddIp(ip string) {
	c.Ips = append(c.Ips, ip)
}

func (c *PingerConfig) SetConsensus(consensus Consensus) {
	c.Consensus = consensus
}

func (c *PingerConfig) SetTimeout(timeout int64, timeoutType TimeoutType) {
	switch timeoutType {
	case SECONDS:
		c.Timeout = time.Second * time.Duration(timeout)
	case MINUTES:
		c.Timeout = time.Minute * time.Duration(timeout)
	}
}

func (c *PingerConfig) GetIps() []string {
	return c.Ips
}

func (c *PingerConfig) Reset() {
	c.Ips = []string{}
}

func (c *PingerConfig) IsAll() bool {
	return ALL == c.Consensus
}

func (c *PingerConfig) IsAny() bool {
	return ANY == c.Consensus
}
