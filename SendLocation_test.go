package telegram

import (
	"fmt"
	"os"
	"testing"
)

func TestSendLocation(t *testing.T) {
	type args struct {
		botToken string
		chatId   string
		location Location
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "TestSendLocation",
			args: args{
				botToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
				chatId:   os.Getenv("TELEGRAM_CHAT_ID"),
				location: Location{
					Longitude: 0,
					Latitude:  0,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SendLocation(tt.args.botToken, tt.args.chatId, tt.args.location)
			if err != tt.want {
				t.Errorf("SendLocation() = %v, want %v", err, tt.want)
			}
		})
	}
}

func ExampleSendLocation() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatId := os.Getenv("TELEGRAM_CHAT_ID")
	location := Location{
		Longitude: 0,
		Latitude:  0,
	}

	message, err := SendLocation(botToken, chatId, location)
	if err != nil {
		panic(err)
	}
	fmt.Println(message.Location)
}
