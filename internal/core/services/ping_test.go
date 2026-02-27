package services

import (
	"context"
	"testing"
)

func TestPingService_Ping(t *testing.T) {
	svc := NewPingService()
	res := svc.Ping(context.Background())
	if res != "pong" {
		t.Errorf("Expected 'pong', got '%s'", res)
	}
}
