package services

import (
	"context"
	"dsnt-challenge/internal/core/domain"
	"dsnt-challenge/internal/core/ports"
)

type echoService struct{}

// NewEchoService creates a new instance of EchoService
func NewEchoService() ports.EchoService {
	return &echoService{}
}

func (s *echoService) Echo(ctx context.Context, req domain.EchoRequest) (domain.EchoResponse, error) {
	return domain.EchoResponse(req), nil
}
