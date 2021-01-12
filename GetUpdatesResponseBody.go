package telegram

// GetUpdatesResponseBody defines the structure of the body returned by the getUpdates method of the telegram API
type GetUpdatesResponseBody struct {
	OK     bool     `json:"ok"`
	Result *Updates `json:"result"`
}
