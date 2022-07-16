package main

import (
	"log"

	"github.co/ted-vo/telegram-bot-message/config"
	"github.co/ted-vo/telegram-bot-message/pkg/bot"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	bot := bot.NewBot(config.Bot)

	_, err = bot.SendMessage("Hello world from bot")
	if err != nil {
		log.Fatal("cannot send message: ", err)
	}
}
