package main

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

type Telegram struct {
	Bot       *gotgbot.Bot
	ChannelID int64
}

func initTelegramBots() *gotgbot.Bot {
	log.Debug("Creating telegram bot")

	TGBot, err := gotgbot.NewBot(config.TelegramToken, nil)
	if err != nil {
		log.Fatal("Error:", "Error init telegram bot:", err.Error())
	}

	return TGBot
}

func NewTelegram() *Telegram {
	return &Telegram{
		Bot:       initTelegramBots(),
		ChannelID: config.TelegramChannelID,
	}
}

func (t *Telegram) SendMessage(message string) (*gotgbot.Message, error) {
	log.Debugf("Sending message to %d", t.ChannelID)

	messageOpts := &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
			IsDisabled: true,
		},
	}
	sentMessage, err := t.Bot.SendMessage(t.ChannelID, message, messageOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	log.Infof("Message successfully sent to chat: %d", t.ChannelID)

	return sentMessage, nil
}

var bot = NewTelegram()
