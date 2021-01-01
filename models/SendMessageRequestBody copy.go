package models

// SendMessageRequestBody defines the request body structure required for the `sendMessage` method for the telegram API
type SendMessageRequestBody struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}
