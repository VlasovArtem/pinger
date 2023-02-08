package webhook

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type Command interface {
	GetShortName() string
	GetDescription() string
	CheckCondition(botPingers *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) error
	OnStart(botPingers *BotPingers, botPinger *BotPinger, message *tgbotapi.Message)
	OnContinue(botPingers *BotPingers, botPinger *BotPinger, message *tgbotapi.Message)
}

type commandFunc = func(botPingers *BotPingers, botPinger *BotPinger, message *tgbotapi.Message)

type commandImpl struct {
	shortName   string
	description string
	onStart     commandFunc
	onContinue  commandFunc
	condition   func(b *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) error
}

func (c *commandImpl) GetShortName() string {
	return c.shortName
}

func (c *commandImpl) GetDescription() string {
	return c.description
}

func (c *commandImpl) CheckCondition(botPingers *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) error {
	if c.condition != nil {
		if err := c.condition(botPingers, botPinger, message); err != nil {
			log.Err(err).Msg("condition error")
			return err
		}
	}
	return nil
}

func (c *commandImpl) OnStart(botPingers *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) {
	if c.onStart != nil {
		c.onStart(botPingers, botPinger, message)
	}
}

func (c *commandImpl) OnContinue(botPingers *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) {
	if c.onContinue != nil {
		c.onContinue(botPingers, botPinger, message)
	}
}

func createCommands(commands ...Command) map[string]Command {
	commandsMap := make(map[string]Command)
	for _, c := range commands {
		commandsMap[c.GetShortName()] = c
	}
	return commandsMap
}

func botIsEnabled(b *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) error {
	if !botPinger.enabled {
		return errors.New("bot is not enabled. Send '/enable' command to enable it")
	}
	return nil
}

func setCurrentCommand(b *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) {
	botPinger.currentCommand = b.commands[message.Text]
}

func unsetCurrentCommand(b *BotPingers, botPinger *BotPinger, message *tgbotapi.Message) {
	botPinger.currentCommand = nil
}
