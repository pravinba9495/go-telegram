package telegram

// GetUserProfilePhotoResponseBody defines the structure of the body returned by the getUserProfilePhotos method of the telegram API
type GetUserProfilePhotoResponseBody struct {
	OK     bool `json:"ok"`
	Result struct {
		Count  int      `json:"total_count"`
		Photos [][]File `json:"photos"`
	} `json:"result"`
}
