package telegram

import (
	"encoding/json"
	"errors"
	"net/http"
)

// GetMe retrieves basic info about the bot, useful for testing bot's auth token
func GetMe(botToken string) (*User, error) {
	response, err := http.Get("https://api.telegram.org/" + botToken + "/getMe")
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("getMe request returned: " + response.Status)
	}
	var result *GetMeResponseBody
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.New("could not convert response to *GetMeResponseBody")
	}
	return result.Result, nil
}
