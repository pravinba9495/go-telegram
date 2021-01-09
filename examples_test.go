package telegram

import (
	"fmt"
	"os"
)

func ExampleGetUpdates() {
	// Bot token generated from BotFather
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Retrieving all updates until the server returns empty response
	var allUpdates []Update
	offset := 0
	for {
		fmt.Println("Using offset: " + fmt.Sprint(offset))
		updates, err := GetUpdates(botToken, fmt.Sprint(offset))
		if err != nil {
			panic(err)
		}
		if len(*updates) > 0 {
			allUpdates = append(allUpdates, *updates...)
			offset = int(allUpdates[len(*updates)-1].UpdateID) + 1
			if len(*updates) < 100 {
				break
			}
		} else {
			break
		}
	}
	fmt.Println("Done !")
	// Output: Using offset: 0
	// Done !
}

func ExampleSendMessage() {
	// Bot token generated from BotFather
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	message := "Hi, I am a message from the telegram bot."
	chatId := os.Getenv("TELEGRAM_CHAT_ID")

	if chatId != "" {
		// Sending a text message
		result, err := SendMessage(botToken, chatId, message)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("[SENT] " + result.Text)
	}
	// Output: [SENT] Hi, I am a message from the telegram bot.
}

func ExampleSetWebhook() {
	// Bot token generated from BotFather
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Setting webhook
	if err := SetWebhook(botToken, ""); err != nil {
		panic(err)
	}
	fmt.Println("Done !")
	// Output: Done !
}

func ExampleSendMedia() {
	// Bot token generated from BotFather
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	chatId := os.Getenv("TELEGRAM_CHAT_ID")
	path := os.Getenv("TELEGRAM_FILE_PATH")

	_, err := SendMedia(botToken, chatId, path)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done !")
	// Output: Done !
}
