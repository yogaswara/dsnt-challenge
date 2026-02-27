package http

import (
	"bytes"
	"context"
	"dsnt-challenge/internal/core/domain"
	"dsnt-challenge/pkg/logger"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	logger.Init()
}

type mockEchoService struct{}

func (m *mockEchoService) Echo(ctx context.Context, req domain.EchoRequest) (domain.EchoResponse, error) {
	if len(req) == 0 {
		return nil, errors.New("request body cannot be empty")
	}
	return domain.EchoResponse(req), nil
}

func TestEchoHandler_Handle(t *testing.T) {
	svc := &mockEchoService{}
	handler := NewEchoHandler(svc)

	// Positive Case
	t.Run("Success POST", func(t *testing.T) {
		payload := []byte(`{"message": "hello test", "age": 25}`)
		req, err := http.NewRequest(http.MethodPost, "/echo", bytes.NewBuffer(payload))
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

		if response["message"] != "hello test" || response["age"] != float64(25) {
			t.Errorf("expected echoed json properly, got: %v", response)
		}
	})

	// Negative Case: Wrong Method
	t.Run("Method GET", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/echo", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.Handle(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})

	// Negative Case: Empty Message
	t.Run("Empty Message Error", func(t *testing.T) {
		payload := []byte(`{}`)
		req, err := http.NewRequest(http.MethodPost, "/echo", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.Handle(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})
}
