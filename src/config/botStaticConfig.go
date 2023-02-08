package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"os"
)

type BotStaticConfig struct {
	Bot            BotConfig
	Chat           Chat
	PingerFilePath string `yaml:"pingerFilePath"`
}

func (c BotStaticConfig) validate() error {
	if err := c.Chat.validate(); err != nil {
		return err
	}
	if c.PingerFilePath == "" {
		return errors.New("PingerFilePath is not set")
	}
	_, err := os.Stat(c.PingerFilePath)
	if err != nil {
		return errors.New("PingerFilePath not exists")
	}
	err = c.Bot.validate()
	if err != nil {
		return err
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
	return &config, config.validate()
}
