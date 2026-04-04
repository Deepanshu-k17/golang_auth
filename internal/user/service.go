package user

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go-auth/internal/auth"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *Repo
	jwtSecret string
}

func NewService(repo *Repo, jwtSecret string) *Service {
	return &Service{repo: repo, jwtSecret: jwtSecret}

}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResult struct {
	Token string     `json:"token"`
	User  PublicUser `json:"user"`
}

func (s *Service) Register(ctx context.Context, input LoginInput) (AuthResult, error) {
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

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResult{}, fmt.Errorf("failed to hash password: %w", err)
	}
	now := time.Now().UTC()

	u := User{
		Email:     email,
		Password:  string(hashBytes),
		Role:      "user",
		CreatedAt: now,
		UpdatedAt: now,
	}
	created, err := s.repo.Create(ctx, &u)
	if err != nil {
		return AuthResult{}, fmt.Errorf("failed to create user: %w", err)
	}

	token, err := auth.CreateToken(s.jwtSecret, created.ID.Hex(), created.Role)
	if err != nil {
		return AuthResult{}, fmt.Errorf("failed to create token: %w", err)
	}
	return AuthResult{Token: token, User: ToPublic(*created)}, nil
}

func (s *Service) Login(ctx context.Context, input LoginInput) (AuthResult, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.TrimSpace(input.Password)
	if email == "" || password == "" {
		return AuthResult{}, errors.New("email and password are required")
	}
	if len(password) < 6 {
		return AuthResult{}, errors.New("password must be at least 6 characters")
	}
	u, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return AuthResult{}, errors.New("invalid email or password")
		}
		return AuthResult{}, fmt.Errorf("failed to find user: %w", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return AuthResult{}, errors.New("invalid email or password")
	}
	token, err := auth.CreateToken(s.jwtSecret, u.ID.Hex(), u.Role)
	if err != nil {
		return AuthResult{}, fmt.Errorf("failed to create token: %w", err)
	}
	return AuthResult{Token: token, User: ToPublic(*u)}, nil

}
