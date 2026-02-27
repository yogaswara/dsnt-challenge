package services

import (
	"context"
	"errors"

	"dsnt-challenge/internal/core/domain"
	"dsnt-challenge/internal/core/ports"
)

type authService struct{}

// NewAuthService creates a new authentication service
func NewAuthService() ports.AuthService {
	return &authService{}
}

func (s *authService) Login(ctx context.Context, req domain.AuthRequest) (domain.AuthResponse, error) {
	if req.Username == "admin" && req.Password == "password" {
		return domain.AuthResponse{Token: "valid-admin-token"}, nil
	}
	return domain.AuthResponse{}, errors.New("invalid credentials")
}

func (s *authService) ValidateToken(ctx context.Context, token string) error {
	if token == "valid-admin-token" {
		return nil
	}
	return errors.New("invalid token")
}
