package files

import "time"

type File struct {
	ID           int       `json:"id"`
	OwnerUserID  int       `json:"owner_user_id"`
	OriginalName string    `json:"original_name"`
	StoragePath  string    `json:"storage_path"`
	MimeType     string    `json:"mime_type"`
	SizeBytes    int64     `json:"size_bytes"`
	Visibility   string    `json:"visibility"`
	CreatedAt    time.Time `json:"created_at"`
}
