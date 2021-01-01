package telegram

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pravinba9495/go-telegram/models"
)

// GetUpdates retrieves incoming updates using long polling
func GetUpdates(botToken string, offset string) (*[]models.Update, error) {
	response, err := http.Get("https://api.telegram.org/" + botToken + "/getUpdates?&allowed_updates=[\"message\"]&offset=" + offset)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("getUpdates request returned: " + response.Status)
	}
	var result *models.GetUpdatesResponseBody
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.New("could not convert response to *models.GetUpdatesBody")
	}
	return result.Result, nil
}
