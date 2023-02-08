package config

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
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

type pingerFileConfig struct {
	Ips       []string
	Consensus string
	Timeout   struct {
		Value int64
		Type  string
	}
}

func (p pingerFileConfig) toPingerConfig() *PingerConfig {
	pingerConfig := PingerConfig{
		Ips:       p.Ips,
		Consensus: Consensus(p.Consensus),
	}
	pingerConfig.SetTimeout(p.Timeout.Value, TimeoutType(p.Timeout.Type))
	return &pingerConfig
}

func (p pingerFileConfig) format() pingerFileConfig {
	p.Consensus = strings.ToLower(p.Consensus)
	p.Timeout.Type = strings.ToLower(p.Timeout.Type)
	return p
}

func (p pingerFileConfig) validate() {
	if p.Timeout.Value <= 0 {
		log.Fatal().Msg("Timeout value must be positive")
	}
	if Consensus(p.Consensus) != ALL && Consensus(p.Consensus) != ANY {
		log.Fatal().Msg("Consensus must be 'all' or 'any'")
	}
	if TimeoutType(p.Timeout.Type) != SECONDS && TimeoutType(p.Timeout.Type) != MINUTES {
		log.Fatal().Msg("Timeout type must be 'seconds' or 'minutes'")
	}
	if len(p.Ips) == 0 {
		log.Fatal().Msg("Ips can't be empty")
	}
}

type PingerConfig struct {
	Ips       []string
	Consensus Consensus
	Timeout   time.Duration
}

func NewPingerConfigFromFile(filePath string) (*PingerConfig, error) {
	if fileContent, err := os.ReadFile(filePath); err != nil {
		return nil, errors.Wrap(err, "Can't read config file")
	} else {
		pingerConfig := pingerFileConfig{}
		err := yaml.Unmarshal(fileContent, &pingerConfig)
		if err != nil {
			return nil, errors.Wrap(err, "Can't parse config file")
		} else {
			formatterPingerConfigFile := pingerConfig.format()
			formatterPingerConfigFile.validate()
			return formatterPingerConfigFile.toPingerConfig(), nil
		}
	}
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
