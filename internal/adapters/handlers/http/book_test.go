package http

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"dsnt-challenge/internal/core/domain"
	"dsnt-challenge/pkg/logger"
)

func init() {
	logger.Init()
}

// mockBooksService is a mock service for testing handler
type mockBooksService struct {
	books map[string]domain.Book
}

func (m *mockBooksService) GetBooks(ctx context.Context, req domain.GetBooksRequest) ([]domain.Book, domain.PaginationMeta, error) {
	return []domain.Book{}, domain.PaginationMeta{}, nil
}

func (m *mockBooksService) GetBookByID(ctx context.Context, id string) (domain.Book, error) {
	if b, ok := m.books[id]; ok {
		return b, nil
	}
	return domain.Book{}, errors.New("book not found")
}

func (m *mockBooksService) CreateBook(ctx context.Context, req domain.CreateBookRequest) (domain.Book, error) {
	if req.Title == "" {
		return domain.Book{}, errors.New("title is required")
	}
	return domain.Book{ID: "1", Title: req.Title, Author: req.Author, Year: req.Year}, nil
}

func (m *mockBooksService) UpdateBook(ctx context.Context, id string, req domain.UpdateBookRequest) (domain.Book, error) {
	if _, ok := m.books[id]; !ok {
		return domain.Book{}, errors.New("book not found")
	}
	return domain.Book{ID: id, Title: req.Title, Author: req.Author, Year: req.Year}, nil
}

func (m *mockBooksService) DeleteBook(ctx context.Context, id string) error {
	if _, ok := m.books[id]; !ok {
		return errors.New("book not found")
	}
	delete(m.books, id)
	return nil
}

func TestBooksHandler(t *testing.T) {
	svc := &mockBooksService{
		books: map[string]domain.Book{
			"1": {ID: "1", Title: "Old Book", Author: "Author", Year: 2010},
		},
	}
	handler := NewBooksHandler(svc)

	t.Run("CreateBook_Success", func(t *testing.T) {
		payload := []byte(`{"title": "New Book", "author": "New Author", "year": 2024}`)
		req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(payload))
		rr := httptest.NewRecorder()

		handler.HandleBooks(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status %v, got %v", http.StatusCreated, rr.Code)
		}
	})

	t.Run("CreateBook_InvalidMethod", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/books", nil)
		rr := httptest.NewRecorder()

		handler.HandleBooks(rr, req)

		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected status %v, got %v", http.StatusMethodNotAllowed, rr.Code)
		}
	})

	t.Run("GetBooks_Success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/books?page=1&limit=10", nil)
		rr := httptest.NewRecorder()

		handler.HandleBooks(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %v, got %v", http.StatusOK, rr.Code)
		}
	})

	t.Run("GetBookByID_Success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/books/1", nil)
		rr := httptest.NewRecorder()

		handler.HandleBookByID(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %v, got %v", http.StatusOK, rr.Code)
		}
	})

	t.Run("GetBookByID_NotFound", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/books/999", nil)
		rr := httptest.NewRecorder()

		handler.HandleBookByID(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status %v, got %v", http.StatusNotFound, rr.Code)
		}
	})

	t.Run("UpdateBook_Success", func(t *testing.T) {
		payload := []byte(`{"title": "Updated Book", "author": "Author", "year": 2024}`)
		req, _ := http.NewRequest(http.MethodPut, "/books/1", bytes.NewBuffer(payload))
		rr := httptest.NewRecorder()

		handler.HandleBookByID(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %v, got %v", http.StatusOK, rr.Code)
		}
	})

	t.Run("DeleteBook_Success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/books/1", nil)
		rr := httptest.NewRecorder()

		handler.HandleBookByID(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %v, got %v", http.StatusOK, rr.Code)
		}
	})
}
