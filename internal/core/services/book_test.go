package services

import (
	"context"
	"dsnt-challenge/internal/core/domain"
	"errors"
	"testing"
)

// mockBookRepo is a simple mock repository for testing
type mockBookRepo struct {
	books map[string]domain.Book
}

func newMockBookRepo() *mockBookRepo {
	return &mockBookRepo{books: make(map[string]domain.Book)}
}

func (m *mockBookRepo) FindAll(ctx context.Context, page, limit int, search string) ([]domain.Book, int, error) {
	var list []domain.Book
	for _, b := range m.books {
		list = append(list, b)
	}
	return list, len(list), nil
}

func (m *mockBookRepo) FindByID(ctx context.Context, id string) (domain.Book, error) {
	if b, ok := m.books[id]; ok {
		return b, nil
	}
	return domain.Book{}, errors.New("book not found")
}

func (m *mockBookRepo) Save(ctx context.Context, book domain.Book) error {
	m.books[book.ID] = book
	return nil
}

func (m *mockBookRepo) Update(ctx context.Context, book domain.Book) error {
	m.books[book.ID] = book
	return nil
}

func (m *mockBookRepo) Delete(ctx context.Context, id string) error {
	delete(m.books, id)
	return nil
}

func TestBooksService(t *testing.T) {
	repo := newMockBookRepo()
	svc := NewBooksService(repo)
	ctx := context.Background()

	t.Run("CreateBook", func(t *testing.T) {
		req := domain.CreateBookRequest{Title: "Mastering Go", Author: "Mihalis", Year: 2021}
		book, err := svc.CreateBook(ctx, req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if book.ID == "" {
			t.Error("expected non-empty ID")
		}
		if book.Title != "Mastering Go" {
			t.Errorf("expected title 'Mastering Go', got '%s'", book.Title)
		}
	})

	t.Run("GetBooks", func(t *testing.T) {
		req := domain.GetBooksRequest{Page: 1, Limit: 10}
		books, meta, err := svc.GetBooks(ctx, req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if meta.TotalItems != 1 {
			t.Errorf("expected 1 item, got %d", meta.TotalItems)
		}
		if len(books) != 1 {
			t.Errorf("expected 1 book in slice, got %d", len(books))
		}
	})

	t.Run("GetBookByID", func(t *testing.T) {
		// First, create to get a valid ID
		created, _ := svc.CreateBook(ctx, domain.CreateBookRequest{Title: "Book 2", Author: "Author 2", Year: 2020})

		found, err := svc.GetBookByID(ctx, created.ID)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if found.Title != "Book 2" {
			t.Errorf("expected title 'Book 2', got '%s'", found.Title)
		}
	})

	t.Run("UpdateBook", func(t *testing.T) {
		created, _ := svc.CreateBook(ctx, domain.CreateBookRequest{Title: "Book 3", Author: "Author 3", Year: 2019})

		updateReq := domain.UpdateBookRequest{Title: "Updated 3", Author: "Updated Author 3", Year: 2020}
		updated, err := svc.UpdateBook(ctx, created.ID, updateReq)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if updated.Title != "Updated 3" {
			t.Errorf("expected title 'Updated 3', got '%s'", updated.Title)
		}
	})

	t.Run("DeleteBook", func(t *testing.T) {
		created, _ := svc.CreateBook(ctx, domain.CreateBookRequest{Title: "Book 4", Author: "Author 4", Year: 2021})

		err := svc.DeleteBook(ctx, created.ID)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		_, err = svc.GetBookByID(ctx, created.ID)
		if err == nil {
			t.Error("expected error 'book not found', got nil")
		}
	})
}
