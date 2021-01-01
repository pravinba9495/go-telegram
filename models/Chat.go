package models

// Chat defines the structure of chat object
type Chat struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}
