package telegram

import (
	"errors"
	"net/http"
	"net/url"
)

// SetWebhook sets the webhook URL for the telegram bot. Setting an empty url will remove the webhook integration
func SetWebhook(botToken string, webhookURL string) error {
	response, err := http.Post("https://api.telegram.org/"+botToken+"/setWebhook?drop_pending_updates=false&allowed_updates=[\"message\"]&url="+url.PathEscape(webhookURL), "application/json", nil)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return errors.New("sendMessage request returned: " + response.Status)
	}
	return nil
}
