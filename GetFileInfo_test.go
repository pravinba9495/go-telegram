package telegram

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestGetFileInfo(t *testing.T) {
	type args struct {
		botToken string
		fileId   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestGetFileInfo",
			args: args{
				botToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
				fileId:   os.Getenv("TELEGRAM_FILE_ID"),
			},
			want:    os.Getenv("TELEGRAM_FILE_ID"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFileInfo(tt.args.botToken, tt.args.fileId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.FileID, tt.want) {
				t.Errorf("GetFileInfo() = %v, want %v", got.FileID, tt.want)
			}
		})
	}
}

func ExampleGetFileInfo() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	fileId := os.Getenv("TELEGRAM_FILE_ID")
	file, err := GetFileInfo(botToken, fileId)
	if err != nil {
		panic(err)
	}
	fmt.Println(file)
}
