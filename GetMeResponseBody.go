package telegram

// GetMeResponseBody defines the structure of the body returned by the getMe method of the telegram API
type GetMeResponseBody struct {
	OK     bool  `json:"ok"`
	Result *User `json:"result"`
}
