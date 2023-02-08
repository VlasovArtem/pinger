package config

import (
	"github.com/pkg/errors"
)

type BotConfig struct {
	Token string
	Url   string
}

func (b *BotConfig) validate() error {
	if b.Token == "" {
		return errors.New("Bot token is not set")
	}
	return nil
}

type Chat struct {
	ChatId   int64 `yaml:"chatId"`
	Username string
}

func (c *Chat) validate() error {
	if c.ChatId == 0 && c.Username == "" {
		return errors.New("ChatId or Username is not set")
	}
	return nil

}
