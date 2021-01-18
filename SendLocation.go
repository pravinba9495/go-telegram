package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// SendLocation sends a location to a recipient
func SendLocation(botToken string, chatId string, location Location) (*Message, error) {
	body := &SendLocationRequestBody{
		ChatID:    chatId,
		Longitude: location.Longitude,
		Latitude:  location.Latitude,
	}
	b, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}
	response, err := http.Post("https://api.telegram.org/"+botToken+"/sendLocation", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("sendLocation request returned: " + response.Status)
	}
	var result *ResponseBody
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.New("could not convert response to *ResponseBody")
	}
	return result.Result, nil
}
