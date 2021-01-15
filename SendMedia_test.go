package telegram

import (
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestSendMedia(t *testing.T) {
	type args struct {
		botToken string
		chatId   string
		path     string
	}
	chatId, _ := strconv.ParseUint(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)
	tests := []struct {
		name    string
		args    args
		want    *Message
		wantErr bool
	}{
		{
			name: "TestSendMedia",
			args: args{
				botToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
				chatId:   os.Getenv("TELEGRAM_CHAT_ID"),
				path:     os.Getenv("TELEGRAM_FILE_PATH"),
			},
			want: &Message{
				Chat: &Chat{
					ID: chatId,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendMedia(tt.args.botToken, tt.args.chatId, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMedia() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Chat.ID, tt.want.Chat.ID) {
				t.Errorf("SendMedia() = %v, want %v", got.Chat.ID, tt.want.Chat.ID)
			}
		})
	}
}
