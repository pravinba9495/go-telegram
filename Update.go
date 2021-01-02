package telegram

type Update struct {
	UpdateID uint64   `json:"update_id"`
	Message  *Message `json:"message"`
}
