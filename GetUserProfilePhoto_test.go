package telegram

import (
	"fmt"
	"os"
	"testing"
)

func TestGetUserProfilePhoto(t *testing.T) {
	type args struct {
		botToken string
		userId   string
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "TestGetUpdates",
			args: args{
				botToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
				userId:   os.Getenv("TELEGRAM_CHAT_ID"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetUserProfilePhoto(tt.args.botToken, tt.args.userId)
			if err != tt.want {
				t.Errorf("GetUserProfilePhoto() = %v, want %v", err, tt.want)
			}
		})
	}
}

func ExampleGetUserProfilePhoto() {
	// Bot token generated from BotFather
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	userId := os.Getenv("TELEGRAM_CHAT_ID")

	files, err := GetUserProfilePhoto(botToken, userId)
	if err != nil {
		panic(err)
	}
	fmt.Print(files)
}
