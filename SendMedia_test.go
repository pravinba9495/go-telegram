package telegram

import (
	"fmt"
	"os"
	"testing"
)

func TestSendMedia(t *testing.T) {
	type args struct {
		botToken string
		chatId   string
		path     string
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "TestSendMedia",
			args: args{
				botToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
				chatId:   os.Getenv("TELEGRAM_CHAT_ID"),
				path:     os.Getenv("TELEGRAM_FILE_PATH"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SendMedia(tt.args.botToken, tt.args.chatId, tt.args.path)
			if err != tt.want {
				t.Errorf("SendMedia() = %v, want %v", err, tt.want)
			}
		})
	}
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
