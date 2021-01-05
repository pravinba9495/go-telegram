package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

// SendMessage sends a text message to a recipient
func SendMessage(botToken string, chatId string, text string) (*Message, error) {
	body := &SendMessageRequestBody{
		ChatID: chatId,
		Text:   text,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	response, err := http.Post("https://api.telegram.org/"+botToken+"/sendMessage", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("sendMessage request returned: " + response.Status)
	}
	var result *SendMessageResponseBody
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.New("could not convert response to *SendMessageResponseBody")
	}
	return result.Result, nil
}

func ExampleSendMessage() {
	// Bot token generated from BotFather
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	message := "Hi, I am a message from the telegram bot."
	chatId := "12345"

	if chatId != "" {
		// Sending a text message
		result, err := SendMessage(botToken, chatId, message)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("[SENT] " + result.Text)
	}
	// Output: [SENT] Hi, I am a message from the telegram bot.
}
