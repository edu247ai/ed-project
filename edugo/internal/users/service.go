package users

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(ctx context.Context, name, email, password, role string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Use PasswordHash (NOT Password)
	u := &User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
		Role:         role,
	}

	// Use Create (NOT CreateUser)
	created, err := s.repo.Create(ctx, email, u.PasswordHash, role)
	if err != nil {
		return nil, err
	}

	// Copy generated fields back
	u.ID = created.ID
	u.CreatedAt = created.CreatedAt

	return u, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (*User, error) {
	u, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare with PasswordHash (NOT Password)
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return nil, errors.New("invalid email or password")
	}

	return u, nil
}

func (s *Service) GetProfile(ctx context.Context, id int) (*User, error) {
	return s.repo.GetByID(ctx, id)
}
