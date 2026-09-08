package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"edu247/edugo/internal/config"
	"edu247/edugo/internal/users"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Service struct {
	cfg  config.Config
	repo *users.Repository
}

func NewService(cfg config.Config, repo *users.Repository) *Service {
	return &Service{cfg: cfg, repo: repo}
}

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func checkPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (s *Service) Register(ctx context.Context, email, password string) (*users.User, string, error) {
	h, err := hashPassword(password)
	if err != nil {
		return nil, "", err
	}

	u, err := s.repo.Create(ctx, email, h, "student")
	if err != nil {
		return nil, "", err
	}

	token, err := GenerateToken(s.cfg.JWTSecret, u.ID, u.Email, u.Role)
	if err != nil {
		return nil, "", err
	}

	return u, token, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (*users.User, string, error) {
	u, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	if err := checkPassword(u.PasswordHash, password); err != nil {
		return nil, "", ErrInvalidCredentials
	}

	token, err := GenerateToken(s.cfg.JWTSecret, u.ID, u.Email, u.Role)
	if err != nil {
		return nil, "", err
	}

	return u, token, nil
}
