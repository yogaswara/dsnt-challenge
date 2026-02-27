package services

import (
	"context"
	"dsnt-challenge/internal/core/ports"
)

type pingService struct{}

// NewPingService creates a new instance of PingService
func NewPingService() ports.PingService {
	return &pingService{}
}

func (s *pingService) Ping(ctx context.Context) string {
	return "pong"
}
