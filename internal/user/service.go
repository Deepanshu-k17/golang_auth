package user

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

type Service struct {
	repo      *Repo
	jwtSecret string
}

func NewService(repo *Repo, jwtSecret string) *Service {
	return &Service{repo: repo, jwtSecret: jwtSecret}

}

type RegisterInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResult struct {
	Token string     `json:"token"`
	User  PublicUser `json:"user"`
}

func (s *Service) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.TrimSpace(input.Password)
	// ... rest of the registration logic
	if email == "" || password == "" {
		return AuthResult{}, errors.New("email and password are required")
	}
	if len(password) < 6 {
		return AuthResult{}, errors.New("password must be at least 6 characters")
	}
	_, err := s.repo.FindByEmail(ctx, email)
	if err == nil {
		return AuthResult{}, errors.New("email already in use")
	}
	if err != nil && !strings.Contains(err.Error(), "user not found") {
		return AuthResult{}, fmt.Errorf("failed to check existing user: %w", err)
	}
}
