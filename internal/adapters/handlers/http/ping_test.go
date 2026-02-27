package http

import (
	"context"
	"dsnt-challenge/pkg/logger"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	logger.Init() // Ensure logger is initialized for tests
}

type mockPingService struct{}

func (m *mockPingService) Ping(ctx context.Context) string {
	return "pong"
}

func TestPingHandler_Handle(t *testing.T) {
	svc := &mockPingService{}
	handler := NewPingHandler(svc)

	// Positive Case
	t.Run("Method GET", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/ping", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.Handle(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var response map[string]interface{}
		json.Unmarshal(rr.Body.Bytes(), &response)

		if response["success"] != true {
			t.Errorf("expected success to be true")
		}
	})

	// Negative Case
	t.Run("Method POST", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/ping", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.Handle(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})
}
