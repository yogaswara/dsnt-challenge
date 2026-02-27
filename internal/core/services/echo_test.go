package services

import (
	"context"
	"dsnt-challenge/internal/core/domain"
	"testing"
)

func TestEchoService_Echo(t *testing.T) {
	svc := NewEchoService()

	// Positive Case
	t.Run("Success", func(t *testing.T) {
		req := domain.EchoRequest{Message: "hello world"}
		res, err := svc.Echo(context.Background(), req)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if res.Message != "hello world" {
			t.Errorf("Expected 'hello world', got '%s'", res.Message)
		}
	})

	// Negative Case
	t.Run("Empty Message Error", func(t *testing.T) {
		req := domain.EchoRequest{Message: ""}
		_, err := svc.Echo(context.Background(), req)
		if err == nil {
			t.Error("Expected error for empty message, got nil")
		}
		if err.Error() != "message cannot be empty" {
			t.Errorf("Expected 'message cannot be empty', got '%s'", err.Error())
		}
	})
}
