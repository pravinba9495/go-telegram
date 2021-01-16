package telegram

import (
	"os"
	"testing"
)

func TestGetFileInfo(t *testing.T) {
	type args struct {
		botToken string
		fileId   string
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "TestGetFileInfo",
			args: args{
				botToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
				fileId:   os.Getenv("TELEGRAM_FILE_ID"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetFileInfo(tt.args.botToken, tt.args.fileId)
			if err != tt.want {
				t.Errorf("GetFileInfo() = %v, want %v", err, tt.want)
			}
		})
	}
}
