package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

// GetUpdates retrieves incoming updates using long polling
func GetUpdates(botToken string, offset string) (*[]Update, error) {
	response, err := http.Get("https://api.telegram.org/" + botToken + "/getUpdates?&allowed_updates=[\"message\"]&offset=" + offset)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("getUpdates request returned: " + response.Status)
	}
	var result *GetUpdatesResponseBody
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.New("could not convert response to *GetUpdatesBody")
	}
	return result.Result, nil
}

func ExampleGetUpdates() {
	// Bot token generated from BotFather
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Retrieving all updates until the server returns empty response
	var allUpdates []Update
	offset := 0
	for {
		log.Println("Using offset: " + fmt.Sprint(offset))
		updates, err := GetUpdates(botToken, fmt.Sprint(offset))
		if err != nil {
			log.Println(err)
			break
		}
		if len(*updates) > 0 {
			allUpdates = append(allUpdates, *updates...)
			offset = int(allUpdates[len(*updates)-1].UpdateID) + 1
		} else {
			break
		}
	}
	for _, update := range allUpdates {
		log.Println("[RECEIVED MESSAGE] " + update.Message.Text)
	}
	// Output: [RECEIVED MESSAGE] Hi, I am a message from the telegram bot.
}
