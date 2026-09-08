package users

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

// ------------------------------------------------------------
// Create user (used by auth.Register)
// ------------------------------------------------------------
func (r *Repository) Create(ctx context.Context, email, passwordHash, role string) (*User, error) {
	u := &User{
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		Name:         "", // optional, can be updated later
	}

	err := r.db.QueryRow(ctx, `
        INSERT INTO users (email, password_hash, role)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `, email, passwordHash, role).Scan(
		&u.ID, &u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}

// ------------------------------------------------------------
// Get user by email (used by auth.Login)
// ------------------------------------------------------------
func (r *Repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	var u User

	err := r.db.QueryRow(ctx, `
        SELECT id, name, email, password_hash, role, created_at
        FROM users
        WHERE email=$1
    `, email).Scan(
		&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.Role, &u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ------------------------------------------------------------
// Get user by ID (used by profile endpoints)
// ------------------------------------------------------------
func (r *Repository) GetByID(ctx context.Context, id int) (*User, error) {
	var u User

	err := r.db.QueryRow(ctx, `
        SELECT id, name, email, role, created_at
        FROM users
        WHERE id=$1
    `, id).Scan(
		&u.ID, &u.Name, &u.Email, &u.Role, &u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
