package telegram

// User defines the structure of user object
type User struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}
