package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// SendMessage sends a text message to a recipient
func SendMessage(botToken string, chatId string, text string) (*Message, error) {
	body := &SendMessageRequestBody{
		ChatID: chatId,
		Text:   text,
	}
	b, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}
	response, err := http.Post("https://api.telegram.org/"+botToken+"/sendMessage", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("sendMessage request returned: " + response.Status)
	}
	var result *SendMessageResponseBody
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.New("could not convert response to *SendMessageResponseBody")
	}
	return result.Result, nil
}
