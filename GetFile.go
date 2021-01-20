package telegram

import (
	"errors"
	"io"
	"net/http"
)

// GetFile retrieves the file by path from telegram API
func GetFile(botToken string, path string) (io.ReadCloser, error) {
	response, err := http.Get("https://api.telegram.org/file/" + botToken + "/" + path)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("getFile request returned: " + response.Status)
	}
	return response.Body, nil
}
