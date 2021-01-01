package models

// Message defines the structure of message object
type Message struct {
	MessageID uint64    `json:"message_id"`
	Chat      *Chat     `json:"chat"`
	Text      string    `json:"text,omitempty"`
	Audio     *File     `json:"audio,omitempty"`
	Document  *File     `json:"document,omitempty"`
	Photo     *[]File   `json:"photo,omitempty"`
	Video     *File     `json:"video,omitempty"`
	Voice     *File     `json:"voice,omitempty"`
	VideoNote *File     `json:"video_note,omitempty"`
	Caption   string    `json:"caption,omitempty"`
	Animation *File     `json:"animation,omitempty"`
	Location  *Location `json:"location,omitempty"`
}
