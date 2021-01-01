package telegram

import "net/http"

// GetFile retrieves basic info about a file and prepares it for downloading
func GetFile(botToken string, fileId string) (*http.Response, error) {
	return http.Get("https://api.telegram.org/" + botToken + "/getFile?file_id=" + fileId)
}
