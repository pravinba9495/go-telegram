package telegram

import (
	"fmt"
	"os"
	"testing"
)

func TestGetUpdates(t *testing.T) {
	type args struct {
		botToken string
		offset   string
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
				offset:   "0",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetUpdates(tt.args.botToken, tt.args.offset)
			if err != tt.want {
				t.Errorf("GetUpdates() = %v, want %v", err, tt.want)
			}
		})
	}
}

func ExampleGetUpdates() {
	// Bot token generated from BotFather
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

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
}
