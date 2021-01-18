package telegram

import (
	"fmt"
	"os"
	"testing"
)

func TestGetFileContentType(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestGetFileContentType",
			args: args{
				path: os.Getenv("TELEGRAM_FILE_PATH"),
			},
			want:    os.Getenv("TELEGRAM_FILE_TYPE"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFileContentType(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileContentType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFileContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleGetFileContentType() {
	path := os.Getenv("TELEGRAM_FILE_PATH")
	mtype, err := GetFileContentType(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(mtype)
}
