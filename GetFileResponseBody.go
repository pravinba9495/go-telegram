package telegram

// GetFileResponseBody defines the structure of the body returned by the getFile method of the telegram API
type GetFileResponseBody struct {
	OK     bool  `json:"ok"`
	Result *File `json:"result"`
}
