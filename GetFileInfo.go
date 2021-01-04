package telegram

import (
	"encoding/json"
	"errors"
	"net/http"
)

// GetFileInfo retrieves basic info about a file and prepares it for downloading
func GetFileInfo(botToken string, fileId string) (*File, error) {
	response, err := http.Get("https://api.telegram.org/" + botToken + "/getFile?file_id=" + fileId)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("getFileInfo request returned: " + response.Status)
	}
	var result *GetFileResponseBody
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.New("could not convert response to *GetFileResponseBody")
	}
	return result.Result, nil
}
