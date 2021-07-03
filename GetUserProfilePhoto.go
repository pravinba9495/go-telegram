package telegram

import (
	"encoding/json"
	"errors"
	"net/http"
)

// GetUserProfilePhoto retrieves profile photos of a user
func GetUserProfilePhoto(botToken string, userId string) ([][]File, error) {
	response, err := http.Get("https://api.telegram.org/" + botToken + "/getUserProfilePhotos?user_id=" + userId + "&limit=1")
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("getUpdates request returned: " + response.Status)
	}
	var result *GetUserProfilePhotoResponseBody
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.New("could not convert response to *GetUpdatesBody")
	}
	return result.Result.Photos, nil
}
