package files

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"edu247/edugo/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	db      *pgxpool.Pool
	storage *storage.Manager
}

func NewService(db *pgxpool.Pool, s *storage.Manager) *Service {
	return &Service{db: db, storage: s}
}

func (s *Service) SaveFile(ctx context.Context, userID int, file multipart.File, header *multipart.FileHeader) (*File, error) {
	ext := filepath.Ext(header.Filename)
	allowed := map[string]bool{
		".pdf":  true,
		".mp3":  true,
		".epub": true,
	}
	if !allowed[ext] {
		return nil, fmt.Errorf("unsupported file type: %s", ext)
	}

	// Build storage path
	storagePath := fmt.Sprintf("uploads/%d_%d%s", userID, time.Now().UnixNano(), ext)

	// Save using multi-storage manager (Local + MinIO + C)
	if err := s.storage.Save(storagePath, file); err != nil {
		return nil, err
	}

	// Insert into DB
	var f File
	err := s.db.QueryRow(ctx, `
        INSERT INTO files (owner_user_id, original_name, storage_path, mime_type, size_bytes, visibility)
        VALUES ($1, $2, $3, $4, $5, 'private')
        RETURNING id, owner_user_id, original_name, storage_path, mime_type, size_bytes, visibility, created_at
    `, userID, header.Filename, storagePath, header.Header.Get("Content-Type"), header.Size).Scan(
		&f.ID, &f.OwnerUserID, &f.OriginalName, &f.StoragePath,
		&f.MimeType, &f.SizeBytes, &f.Visibility, &f.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &f, nil
}
