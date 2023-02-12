package config

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
)

type BotStaticConfig struct {
	Bot              BotConfig
	MaxNumberOfChats int    `yaml:"maxNumberOfChats"`
	Token            string `yaml:"token"`
}

func (c BotStaticConfig) validate() error {
	err := c.Bot.validate()
	if err != nil {
		return err
	}
	if c.MaxNumberOfChats <= 0 {
		return errors.New("MaxNumberOfChats is not set")
	}
	if c.Token == "" {
		return errors.New("Token is not set")
	}
	return nil
}

func NewBotStaticConfigFromFile(filepath string) (*BotStaticConfig, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "Error while reading config file")
	}

	config := BotStaticConfig{}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, errors.Wrap(err, "Error while parsing config file")
	}
	if config.Token == "" {
		config.Token = uuid.New().String()
		log.Info().Msg("Token is not set, generated new one: " + config.Token)
	}
	return &config, config.validate()
}
