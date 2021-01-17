package telegram

// Update defines the structure of update object
type Update struct {
	UpdateID uint64   `json:"update_id"`
	Message  *Message `json:"message"`
}

// Updates defines an array of updates
type Updates []Update
