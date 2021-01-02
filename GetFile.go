package telegram

import (
	"errors"
	"io"
	"net/http"
)

// GetFile retrieves basic info about a file and prepares it for downloading
func GetFile(botToken string, fileId string) (io.ReadCloser, error) {
	response, err := http.Get("https://api.telegram.org/" + botToken + "/getFile?file_id=" + fileId)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("getFile request returned: " + response.Status)
	}
	return response.Body, nil
}
