# go-telegram
[![build](https://github.com/pravinba9495/go-telegram/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/pravinba9495/go-telegram/actions/workflows/build.yml) ![](https://img.shields.io/github/license/pravinba9495/go-telegram) ![](https://goreportcard.com/badge/github.com/pravinba9495/go-telegram) ![](https://godoc.org/github.com/pravinba9495/go-telegram?status.svg) ![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/pravinba9495/go-telegram?include_prereleases)

## Introduction
Send and receive messages using the Telegram Bot API with ease. Supports plain text messages, media attachments and location data, written in Go.

## Features
- Send and receive text messages
- Send and receive media files
- Update webhook URL

## Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/pravinba9495/go-telegram"
)

func main() {
	// Bot token generated from BotFather
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	text := "Hi, I am a message from the telegram bot."
	chatId := os.Getenv("TELEGRAM_CHAT_ID")
	path := os.Getenv("TELEGRAM_FILE_PATH")

	// Retrieving all updates until the server returns empty response
	var allUpdates Updates
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

	for _, update := range allUpdates {
		fmt.Println("[RECEIVED] " + update.Message.Text)
	}

        // Sending a text message
	result, err := telegram.SendMessage(botToken, chatId, text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("[SENT] " + result.Text)

        // Sending a media attachment
	message, err := telegram.SendMedia(botToken, chatId, path)
	if err != nil {
		panic(err)
	}

	fmt.Println(message.Text)

	// Setting webhook
	if err := telegram.SetWebhook(botToken, "https://example.com/webhook"); err != nil {
		panic(err)
	}
}

```

## Documentation
The complete documentation is available at https://pkg.go.dev/github.com/pravinba9495/go-telegram

## License
MIT
