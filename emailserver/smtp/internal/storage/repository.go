package storage

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) InsertMessage(ctx context.Context, msg *Message) error {
	return r.insertMessage(ctx, r.db, msg)
}

func (r *Repository) insertMessage(ctx context.Context, db interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
}, msg *Message) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO messages (id, sender, body, created_at)
         VALUES ($1, $2, $3, $4)`,
		msg.ID, msg.Sender, msg.Body, msg.CreatedAt,
	)
	return err
}

func (r *Repository) InsertEnvelope(ctx context.Context, env *Envelope) error {
	return r.insertEnvelope(ctx, r.db, env)
}

func (r *Repository) insertEnvelope(ctx context.Context, db interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
}, env *Envelope) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO envelopes (id, message_id, helo, mail_from, created_at)
         VALUES ($1, $2, $3, $4, $5)`,
		env.ID, env.MessageID, env.Helo, env.MailFrom, env.CreatedAt,
	)
	return err
}

func (r *Repository) InsertRecipient(ctx context.Context, rcpt *Recipient) error {
	return r.insertRecipient(ctx, r.db, rcpt)
}

func (r *Repository) insertRecipient(ctx context.Context, db interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
}, rcpt *Recipient) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO recipients (id, envelope_id, address, created_at)
         VALUES ($1, $2, $3, $4)`,
		rcpt.ID, rcpt.EnvelopeID, rcpt.Address, rcpt.CreatedAt,
	)
	return err
}
