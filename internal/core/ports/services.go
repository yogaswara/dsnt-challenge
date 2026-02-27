package ports

import (
	"context"
	"dsnt-challenge/internal/core/domain"
)

// PingService defines the interface for ping related operations
type PingService interface {
	Ping(ctx context.Context) string
}

// EchoService defines the interface for echo related operations
type EchoService interface {
	Echo(ctx context.Context, req domain.EchoRequest) (domain.EchoResponse, error)
}
