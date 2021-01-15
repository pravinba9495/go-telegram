package telegram

import (
	"os"
	"testing"
)

func TestSetWebhook(t *testing.T) {
	type args struct {
		botToken   string
		webhookURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Webhook",
			args: args{
				botToken:   os.Getenv("TELEGRAM_BOT_TOKEN"),
				webhookURL: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetWebhook(tt.args.botToken, tt.args.webhookURL); (err != nil) != tt.wantErr {
				t.Errorf("SetWebhook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
