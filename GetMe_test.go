package telegram

import (
	"os"
	"reflect"
	"testing"
)

func TestGetMe(t *testing.T) {
	type args struct {
		botToken string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "TestGetMe",
			args: args{
				botToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
			},
			want: &User{
				ID: 1827965981,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMe(tt.args.botToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("GetMe() = %v, want %v", got.ID, tt.want.ID)
			}
		})
	}
}
