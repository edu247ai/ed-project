package chatbridge

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveMessage(ctx context.Context, m *ChatMessage) error {
	return r.db.QueryRow(ctx, `
        INSERT INTO chat_messages (room_id, user_id, content)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `, m.RoomID, m.UserID, m.Content).Scan(&m.ID, &m.CreatedAt)
}

func (r *Repository) ListMessages(ctx context.Context, roomID int) ([]ChatMessage, error) {
	rows, err := r.db.Query(ctx, `
        SELECT id, room_id, user_id, content, created_at
        FROM chat_messages
        WHERE room_id=$1
        ORDER BY created_at ASC
    `, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []ChatMessage
	for rows.Next() {
		var m ChatMessage
		rows.Scan(&m.ID, &m.RoomID, &m.UserID, &m.Content, &m.CreatedAt)
		list = append(list, m)
	}
	return list, nil
}
