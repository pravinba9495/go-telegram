package telegram

// Chat defines the structure of chat object
type Chat struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
