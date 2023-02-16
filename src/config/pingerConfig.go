package config

import (
	"time"
)

type Quorum string

const (
	ALL Quorum = "all"
	ANY Quorum = "any"
)

type TimeoutType string

const (
	SECONDS TimeoutType = "seconds"
	MINUTES TimeoutType = "minutes"
)

type PingerConfig struct {
	Ips     []string
	Quorum  Quorum
	Timeout time.Duration
}

func NewConfig(consensus Quorum, timeout time.Duration) *PingerConfig {
	return &PingerConfig{
		Ips:     []string{},
		Quorum:  consensus,
		Timeout: timeout,
	}
}

func (c *PingerConfig) AddIp(ip string) {
	c.Ips = append(c.Ips, ip)
}

func (c *PingerConfig) SetQuorum(quorum Quorum) {
	c.Quorum = quorum
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
	return ALL == c.Quorum
}

func (c *PingerConfig) IsAny() bool {
	return ANY == c.Quorum
}
