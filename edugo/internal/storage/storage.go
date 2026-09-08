package storage

import (
	"io"
	"time"
)

// Storage defines the unified interface for all storage backends.
// Any storage layer (Local, MinIO, C) must implement all methods below.
type Storage interface {
	// Save writes a file to the storage backend.
	Save(path string, data io.Reader) error

	// Get retrieves a file as a stream.
	Get(path string) (io.ReadCloser, error)

	// Delete removes a file from the backend.
	Delete(path string) error

	// PresignURL returns a temporary public URL for downloading the file.
	// Local and C storage may return an error since they do not support presigned URLs.
	PresignURL(path string, expiry time.Duration) (string, error)
}
