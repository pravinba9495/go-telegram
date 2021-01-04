package telegram

import (
	"errors"
	"io"
	"net/http"
)

// GetFile retrieves the file by filePath from telegram API
func GetFile(botToken string, filePath string) (io.ReadCloser, error) {
	response, err := http.Get("https://api.telegram.org/file/" + botToken + "/" + filePath)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("getFile request returned: " + response.Status)
	}
	return response.Body, nil
}
