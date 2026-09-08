package storage

import (
	"fmt"
	"io"
	"os"
	"time"
)

type LocalStorage struct {
	BasePath string
}

func NewLocal(base string) *LocalStorage {
	return &LocalStorage{BasePath: base}
}

func (l *LocalStorage) Save(path string, data io.Reader) error {
	full := l.BasePath + "/" + path

	// Ensure directory exists
	if err := os.MkdirAll(l.BasePath, 0755); err != nil {
		return err
	}

	f, err := os.Create(full)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, data)
	return err
}

func (l *LocalStorage) Get(path string) (io.ReadCloser, error) {
	full := l.BasePath + "/" + path
	return os.Open(full)
}

func (l *LocalStorage) Delete(path string) error {
	full := l.BasePath + "/" + path
	return os.Remove(full)
}

// ------------------------------------------------------------
// REQUIRED BY Storage INTERFACE
// ------------------------------------------------------------
func (l *LocalStorage) PresignURL(path string, expiry time.Duration) (string, error) {
	// Local storage cannot generate presigned URLs
	return "", fmt.Errorf("local storage does not support presigned URLs")
}
