package models

// GetUpdatesResponseBody defines the structure of the body returned by the getUpdates method of the telegram API
type GetUpdatesResponseBody struct {
	OK     bool      `json:"ok"`
	Result *[]Update `json:"result"`
}
