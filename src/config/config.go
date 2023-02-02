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

type Config struct {
	Ips       []string
	Consensus Consensus
	Timeout   time.Duration
}

func NewConfig(consensus Consensus, timeout time.Duration) *Config {
	return &Config{
		Ips:       []string{},
		Consensus: consensus,
		Timeout:   timeout,
	}
}

func (c *Config) AddIp(ip string) {
	c.Ips = append(c.Ips, ip)
}

func (c *Config) SetConsensus(consensus Consensus) {
	c.Consensus = consensus
}

func (c *Config) SetTimeout(timeout int64, timeoutType TimeoutType) {
	switch timeoutType {
	case SECONDS:
		c.Timeout = time.Second * time.Duration(timeout)
	case MINUTES:
		c.Timeout = time.Minute * time.Duration(timeout)
	}
}

func (c *Config) GetIps() []string {
	return c.Ips
}

func (c *Config) Reset() {
	c.Ips = []string{}
}

func (c *Config) IsAll() bool {
	return ALL == c.Consensus
}

func (c *Config) IsAny() bool {
	return ANY == c.Consensus
}
