package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"email-server/smtp/internal/queue/pkg"
)

type PostgresQueueStorage struct {
	db *sql.DB
}

func NewPostgresQueueStorage(db *sql.DB) *PostgresQueueStorage {
	return &PostgresQueueStorage{db: db}
}

func (s *PostgresQueueStorage) LoadQueueItems(ctx context.Context) ([]pkg.QueueItem, error) {
	rows, err := s.db.QueryContext(ctx,
		`WITH due AS (
			SELECT id FROM queue_items
			WHERE next_attempt <= CURRENT_TIMESTAMP
			  AND (lease_until IS NULL OR lease_until < CURRENT_TIMESTAMP)
			ORDER BY next_attempt
			FOR UPDATE SKIP LOCKED
			LIMIT 100
		)
		UPDATE queue_items AS item
		SET lease_until = CURRENT_TIMESTAMP + INTERVAL '15 minutes'
		FROM due
		WHERE item.id = due.id
		RETURNING item.id, item.message_id, item.recipient, item.attempts, item.next_attempt`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]pkg.QueueItem, 0)
	for rows.Next() {
		var item pkg.QueueItem
		if err := rows.Scan(&item.ID, &item.MessageID, &item.Recipient, &item.Attempts, &item.NextAttempt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (s *PostgresQueueStorage) LoadMessage(ctx context.Context, messageID string) ([]byte, string, error) {
	var body []byte
	var sender string
	err := s.db.QueryRowContext(ctx,
		`SELECT body, sender FROM messages WHERE id = $1`, messageID,
	).Scan(&body, &sender)
	if err != nil {
		return nil, "", fmt.Errorf("load queued message: %w", err)
	}
	return body, sender, nil
}

func (s *PostgresQueueStorage) ScheduleRetry(ctx context.Context, id string, next time.Time, reason string) error {
	_, err := s.db.ExecContext(ctx,
		`UPDATE queue_items SET attempts = attempts + 1, next_attempt = $2, last_error = $3, lease_until = NULL WHERE id = $1`,
		id, next, reason,
	)
	return err
}

func (s *PostgresQueueStorage) DeleteQueueItem(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM queue_items WHERE id = $1`, id)
	return err
}
