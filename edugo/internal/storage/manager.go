package storage

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

// Manager coordinates multi‑layer storage:
// Local → MinIO → C
//
// Save: writes to all layers
// Get:  prefers MinIO → C → Local
// Delete: removes from all layers
// PresignURL: only MinIO supports it
type Manager struct {
	Local Storage
	Minio Storage
	C     Storage
}

// ------------------------------------------------------------
// SAVE — write file to all storage layers
// ------------------------------------------------------------
func (m *Manager) Save(path string, data io.Reader) error {
	// Read file once into memory
	buf, err := io.ReadAll(data)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Create new readers for each backend
	localReader := bytes.NewReader(buf)
	minioReader := bytes.NewReader(buf)
	cReader := bytes.NewReader(buf)

	// Save to Local
	if err := m.Local.Save(path, localReader); err != nil {
		return fmt.Errorf("local save failed: %w", err)
	}

	// Save to MinIO
	if err := m.Minio.Save(path, minioReader); err != nil {
		return fmt.Errorf("minio save failed: %w", err)
	}

	// Save to C storage
	if err := m.C.Save(path, cReader); err != nil {
		return fmt.Errorf("c storage save failed: %w", err)
	}

	return nil
}

// ------------------------------------------------------------
// GET — prefer MinIO → C → Local
// ------------------------------------------------------------
func (m *Manager) Get(path string) (io.ReadCloser, error) {
	// Try MinIO first
	if r, err := m.Minio.Get(path); err == nil {
		return r, nil
	}

	// Try C storage
	if r, err := m.C.Get(path); err == nil {
		return r, nil
	}

	// Fallback to Local
	return m.Local.Get(path)
}

// ------------------------------------------------------------
// DELETE — remove from all layers
// ------------------------------------------------------------
func (m *Manager) Delete(path string) error {
	// Delete everywhere (ignore errors)
	_ = m.Local.Delete(path)
	_ = m.Minio.Delete(path)
	_ = m.C.Delete(path)

	return nil
}

// ------------------------------------------------------------
// PRESIGN URL — only MinIO supports it
// ------------------------------------------------------------
func (m *Manager) PresignURL(path string, expiry time.Duration) (string, error) {
	if m.Minio == nil {
		return "", fmt.Errorf("minio storage not configured")
	}

	// Type assert to MinioStorage
	if ms, ok := m.Minio.(*MinioStorage); ok {
		return ms.PresignURL(path, expiry)
	}

	return "", fmt.Errorf("presigned URLs only supported on MinIO")
}
