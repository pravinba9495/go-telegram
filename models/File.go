package models

// File defines the structure of file object
type File struct {
	FileID   string `json:"file_id"`
	FilePath string `json:"file_path,omitempty"`
}
