package services

import (
	"context"
	"dsnt-challenge/internal/core/domain"
	"encoding/json"
	"reflect"
	"testing"
)

func TestEchoService_Echo(t *testing.T) {
	svc := NewEchoService()

	// Positive Case
	t.Run("Success Dictionary", func(t *testing.T) {
		req := domain.EchoRequest(json.RawMessage(`{"message": "hello world", "foo": "bar"}`))
		res, err := svc.Echo(context.Background(), req)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(req, domain.EchoRequest(res)) {
			t.Errorf("Expected correct echoed values, got %s", string(res))
		}
	})

	// Positive Case
	t.Run("Empty Map", func(t *testing.T) {
		req := domain.EchoRequest(json.RawMessage(`{}`))
		res, err := svc.Echo(context.Background(), req)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(req, domain.EchoRequest(res)) {
			t.Errorf("Expected correct echoed values, got %s", string(res))
		}
	})
}
