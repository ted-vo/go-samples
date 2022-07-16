package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.co/ted-vo/telegram-bot-message/config"
)

type Bot struct {
	token   string
	chat_id string
}

func NewBot(config config.BotConfig) Bot {
	return Bot{
		token:   config.Token,
		chat_id: config.ChatId,
	}
}

func (bot *Bot) getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", bot.token)
}

func (bot *Bot) GetUpdates() {

}

func (bot *Bot) SendMessage(text string) (bool, error) {

	url := fmt.Sprintf("%s/sendMessage", bot.getUrl())
	body, _ := json.Marshal(map[string]string{
		"chat_id": bot.chat_id,
		"text":    text,
	})
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	log.Println(fmt.Sprintf("Message: '%s' was sent", text))
	log.Println(fmt.Sprintf("Response JSON: %s", string(body)))

	return true, nil
}
