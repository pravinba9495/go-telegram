package telegram

import (
	"encoding/json"
	"errors"
	"net/http"
)

// GetUpdates retrieves incoming updates using long polling
func GetUpdates(botToken string, offset string) (*[]Update, error) {
	response, err := http.Get("https://api.telegram.org/" + botToken + "/getUpdates?&allowed_updates=[\"message\"]&offset=" + offset)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("getUpdates request returned: " + response.Status)
	}
	var result *GetUpdatesResponseBody
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.New("could not convert response to *GetUpdatesBody")
	}
	return result.Result, nil
}
