package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pravinba9495/go-telegram/models"
	"github.com/pravinba9495/go-telegram/telegram"
)

func main() {

	// Bot token generated from BotFather
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	message := "Hi, I am a message from the telegram bot."

	var chatId string

	// Setting webhook
	if err := telegram.SetWebhook(botToken, ""); err != nil {
		log.Println(err)
		return
	}

	// Retrieving all updates until the server returns empty response
	var allUpdates []models.Update
	offset := 0
	for {
		log.Println("Using offset: " + fmt.Sprint(offset))
		updates, err := telegram.GetUpdates(botToken, fmt.Sprint(offset))
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
		chatId = fmt.Sprint(update.Message.Chat.ID)
	}

	if chatId != "" {
		// Sending a text message
		result, err := telegram.SendMessage(botToken, chatId, message)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("[SENT] " + result.Text)
	}
}
