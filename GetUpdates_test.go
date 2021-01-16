package telegram

import (
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
