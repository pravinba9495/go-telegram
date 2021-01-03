package telegram

import (
	"fmt"
	"log"
	"os"
)

// Set webhook
func Example_setWebhook() {
	// Bot token generated from BotFather
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Setting webhook
	if err := SetWebhook(botToken, ""); err != nil {
		log.Println(err)
	}
}

// Get updates
func Example_getUpdates() {
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
}

// Send a text message
func Example_sendMessage() {
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
}
