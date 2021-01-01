package models

// SendMessageResponseBody defines the structure of the body returned by the `sendMessage` method of the telegram API
type SendMessageResponseBody struct {
	OK     bool     `json:"ok"`
	Result *Message `json:"result"`
}
