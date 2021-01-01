package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pravinba9495/go-telegram/models"
)

// SendMessage sends a text message
func SendMessage(botToken string, chatId string, text string) (*models.Message, error) {
	body := &models.SendMessageRequestBody{
		ChatID: chatId,
		Text:   text,
	}
	b, err := json.Marshal(body)
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
	var result *models.SendMessageResponseBody
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.New("could not convert response to *models.GetUpdatesBody")
	}
	return result.Result, nil
}
