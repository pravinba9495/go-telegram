package telegram

import (
	"os"
	"reflect"
	"testing"
)

func TestSendMessage(t *testing.T) {
	type args struct {
		botToken string
		chatId   string
		text     string
	}
	tests := []struct {
		name    string
		args    args
		want    *Message
		wantErr bool
	}{
		{
			name: "TestSendMessage",
			args: args{
				botToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
				chatId:   os.Getenv("TELEGRAM_CHAT_ID"),
				text:     "Hello World",
			},
			want: &Message{
				Text: "Hello World",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendMessage(tt.args.botToken, tt.args.chatId, tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Text, tt.want.Text) {
				t.Errorf("SendMessage() = %v, want %v", got.Text, tt.want.Text)
			}
		})
	}
}
