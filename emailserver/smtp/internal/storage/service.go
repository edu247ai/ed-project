package storage

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type StorageService struct {
	repo *Repository
	db   *sql.DB
}

func NewStorageService(repo *Repository) *StorageService {
	return &StorageService{repo: repo}
}

func NewPostgresStorage(dsn string) (*StorageService, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("connect to postgres: %w", err)
	}
	if _, err := db.ExecContext(ctx, schema); err != nil {
		db.Close()
		return nil, fmt.Errorf("apply smtp schema: %w", err)
	}
	return &StorageService{repo: NewRepository(db), db: db}, nil
}

func (s *StorageService) Close() error {
	if s.db == nil {
		return nil
	}
	return s.db.Close()
}

func (s *StorageService) SaveIncomingMessage(ctx context.Context, helo, mailFrom string, rcpts []string, body []byte) (string, error) {
	if s.db == nil {
		return "", fmt.Errorf("postgres storage is not configured")
	}
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	msgID, err := NewUUID()
	if err != nil {
		return "", err
	}
	msg := &Message{
		ID:        msgID,
		Sender:    mailFrom,
		Body:      body,
		CreatedAt: time.Now(),
	}

	if err := s.repo.insertMessage(ctx, tx, msg); err != nil {
		return "", err
	}

	envID, err := NewUUID()
	if err != nil {
		return "", err
	}
	env := &Envelope{
		ID:        envID,
		MessageID: msg.ID,
		Helo:      helo,
		MailFrom:  mailFrom,
		CreatedAt: time.Now(),
	}

	if err := s.repo.insertEnvelope(ctx, tx, env); err != nil {
		return "", err
	}

	for _, r := range rcpts {
		recipientID, err := NewUUID()
		if err != nil {
			return "", err
		}
		queueID, err := NewUUID()
		if err != nil {
			return "", err
		}
		rcpt := &Recipient{
			ID:         recipientID,
			EnvelopeID: env.ID,
			Address:    r,
			CreatedAt:  time.Now(),
		}
		if err := s.repo.insertRecipient(ctx, tx, rcpt); err != nil {
			return "", err
		}
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO queue_items (id, message_id, recipient, attempts, next_attempt, created_at)
			 VALUES ($1, $2, $3, 0, $4, $4)`,
			queueID, msg.ID, r, time.Now(),
		); err != nil {
			return "", err
		}
	}
	if err := tx.Commit(); err != nil {
		return "", err
	}

	return msg.ID, nil
}

type MemoryStorage struct {
	mu       sync.Mutex
	messages []Message
	queue    []QueueItem
}

type QueueItem struct {
	MessageID string
	Recipient string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) SaveIncomingMessage(_ context.Context, helo, mailFrom string, rcpts []string, body []byte) (string, error) {
	if len(rcpts) == 0 {
		return "", fmt.Errorf("at least one recipient is required")
	}
	id, err := NewUUID()
	if err != nil {
		return "", err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messages = append(s.messages, Message{ID: id, Sender: mailFrom, Body: append([]byte(nil), body...), CreatedAt: time.Now()})
	for _, recipient := range rcpts {
		s.queue = append(s.queue, QueueItem{MessageID: id, Recipient: recipient})
	}
	return id, nil
}

func (s *MemoryStorage) MessageCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.messages)
}

func (s *MemoryStorage) QueueCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.queue)
}

//go:embed migrations.sql
var schema string
