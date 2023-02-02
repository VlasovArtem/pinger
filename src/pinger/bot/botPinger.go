package bot

import (
	"fmt"
	"github.com/VlasovArtem/pinger/src/pinger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/maps"
	"os"
	"strconv"
)

const BOT_TOKEN string = "TELEGRAM_BOT_TOKEN"
const BOT_URL string = "TELEGRAM_BOT_URL"
const ENABLING_TOKEN string = "TELEGRAM_BOT_ENABLING_TOKEN"

type BotPinger struct {
	*pinger.Pinger
	chatId         int64
	enabled        bool
	currentCommand Command
}

type BotPingers struct {
	enablingToken string
	botApi        *tgbotapi.BotAPI
	pingers       map[int64]*BotPinger
	commands      map[string]Command
}

func NewBotPingers() *BotPingers {
	enablingToken := os.Getenv(ENABLING_TOKEN)
	if enablingToken == "" {
		log.Fatal().Msg("enabling token is not set")
	}
	botToken := os.Getenv(BOT_TOKEN)
	if botToken == "" {
		log.Fatal().Msg("bot token is not set")
	}
	botUrl := os.Getenv(BOT_URL)
	if botUrl == "" {
		botUrl = tgbotapi.APIEndpoint
	}
	botApi, err := tgbotapi.NewBotAPIWithAPIEndpoint(botToken, botUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("bot api init error")
	}

	baseCommands := createCommands(
		NewEnableCommand(),
		NewHelpCommand(),
		NewStartCommand(),
		NewStopCommand(),
		NewSettingsCommand(),
		NewStatusCommand(),
		NewConsensusCommand(),
		NewIpCommand(),
		NewResetCommand(),
	)

	timeoutCommands := createCommands(NewTimeoutCommands()...)
	maps.Copy(baseCommands, timeoutCommands)

	return &BotPingers{
		enablingToken: enablingToken,
		botApi:        botApi,
		pingers:       make(map[int64]*BotPinger),
		commands:      baseCommands,
	}
}

func (b *BotPingers) PerformUpdate(update tgbotapi.Update) {
	b.performPingerChatOperation(update)
	b.performUpdateOnMessage(update.Message)
}

func (b *BotPingers) performPingerChatOperation(update tgbotapi.Update) {
	if myChatMember := update.MyChatMember; myChatMember != nil {
		b.performUpdateOnChat(*myChatMember)
	} else if message := update.Message; message != nil {
		chatID := message.Chat.ID
		if _, ok := b.pingers[chatID]; !ok {
			log.Info().Msg("Pinger is not registered for chat with id: " + strconv.FormatInt(chatID, 10) + ". Adding new pinger")
			b.addNewPinger(chatID)
		}
	}
}

func (b *BotPingers) performUpdateOnChat(member tgbotapi.ChatMemberUpdated) {
	if member.NewChatMember.HasLeft() {
		delete(b.pingers, member.Chat.ID)
	} else if member.NewChatMember.IsMember {
		b.addNewPinger(member.Chat.ID)
	}
}

func (b *BotPingers) addNewPinger(id int64) {
	b.pingers[id] = &BotPinger{
		Pinger: pinger.NewPinger(
			pinger.NewBotProcessor(&pinger.LightBotFormatter{}, b.botApi, id),
			pinger.NewDefaultPingProvider(),
		),
		chatId: id,
	}
}

func (b *BotPingers) performUpdateOnMessage(message *tgbotapi.Message) {
	if message == nil {
		return
	}
	botPinger := b.pingers[message.Chat.ID]

	for _, entity := range message.Entities {
		if entity.IsCommand() {
			if botPinger.currentCommand != nil {
				b.sendMessage(
					botPinger.chatId,
					fmt.Sprintf(
						"Please finish current command ('%s') before starting new one",
						botPinger.currentCommand.GetShortName(),
					),
				)
				return
			} else {
				b.performCommand(b.commands[message.Text], *message)
				return
			}
		}
	}

	if botPinger.currentCommand != nil {
		botPinger.currentCommand.OnContinue(b, botPinger, message)
	}
}

func (b *BotPingers) performCommand(command Command, message tgbotapi.Message) {
	if command == nil {
		b.sendMessage(
			message.Chat.ID,
			"Unknown command "+message.Text,
		)
		return
	}

	botPinger := b.pingers[message.Chat.ID]

	if err := command.CheckCondition(b, botPinger, &message); err != nil {
		b.sendMessage(
			message.Chat.ID,
			err.Error(),
		)
		return
	}

	command.OnStart(b, botPinger, &message)
}

func (b *BotPingers) sendMessage(chatId int64, text string) tgbotapi.Message {
	msg, err := b.botApi.Send(tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: chatId,
		},
		Text: text,
	})
	if err != nil {
		log.Err(err).Msg("send message error")
	}
	return msg
}
