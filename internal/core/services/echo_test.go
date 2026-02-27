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
		req := domain.EchoRequest{"message": "hello world", "foo": "bar"}
		res, err := svc.Echo(context.Background(), req)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if res["message"] != "hello world" || res["foo"] != "bar" {
			t.Errorf("Expected correct echoed values, got %v", res)
		}
	})

	// Negative Case
	t.Run("Empty Message Error", func(t *testing.T) {
		req := domain.EchoRequest{}
		_, err := svc.Echo(context.Background(), req)
		if err == nil {
			t.Error("Expected error for empty message, got nil")
		}
		if err.Error() != "request body cannot be empty" {
			t.Errorf("Expected 'request body cannot be empty', got '%s'", err.Error())
		}
	})
}
