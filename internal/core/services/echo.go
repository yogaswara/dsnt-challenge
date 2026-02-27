package services

import (
	"context"
	"dsnt-challenge/internal/core/domain"
	"dsnt-challenge/internal/core/ports"
	"errors"
)

type echoService struct{}

// NewEchoService creates a new instance of EchoService
func NewEchoService() ports.EchoService {
	return &echoService{}
}

func (s *echoService) Echo(ctx context.Context, req domain.EchoRequest) (domain.EchoResponse, error) {
	if req.Message == "" {
		return domain.EchoResponse{}, errors.New("message cannot be empty")
	}

	return domain.EchoResponse{
		Message: req.Message,
	}, nil
}
