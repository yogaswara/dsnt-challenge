package ports

import (
	"context"

	"dsnt-challenge/internal/core/domain"
)

// AuthService handles authentication logic
type AuthService interface {
	Login(ctx context.Context, req domain.AuthRequest) (domain.AuthResponse, error)
	ValidateToken(ctx context.Context, token string) error
}
