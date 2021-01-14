package telegram

// SendLocationRequestBody defines the request body structure required for the `sendLocation` method for the telegram API
type SendLocationRequestBody struct {
	ChatID    string  `json:"chat_id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
