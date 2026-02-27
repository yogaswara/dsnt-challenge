package memory

import (
	"context"
	"dsnt-challenge/internal/core/domain"
	"testing"
)

func TestBookRepository(t *testing.T) {
	repo := NewBookRepository()
	ctx := context.Background()

	// 1. Test Save and FindByID
	t.Run("Save and FindByID", func(t *testing.T) {
		book := domain.Book{ID: "1", Title: "Go Book", Author: "John", Year: 2020}
		err := repo.Save(ctx, book)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		found, err := repo.FindByID(ctx, "1")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if found.Title != "Go Book" {
			t.Errorf("expected title 'Go Book', got '%s'", found.Title)
		}
	})

	// 2. Test FindAll with Pagination
	t.Run("FindAll Pagination", func(t *testing.T) {
		repo.Save(ctx, domain.Book{ID: "2", Title: "Book 2", Author: "Jane", Year: 2021})
		repo.Save(ctx, domain.Book{ID: "3", Title: "Book 3", Author: "Doe", Year: 2022})

		// We have 3 books now: "1", "2", "3"
		books, total, err := repo.FindAll(ctx, 1, 2, "")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if total != 3 {
			t.Errorf("expected total 3, got %d", total)
		}
		if len(books) != 2 {
			t.Errorf("expected 2 books, got %d", len(books))
		}

		booksPage2, _, _ := repo.FindAll(ctx, 2, 2, "")
		if len(booksPage2) != 1 {
			t.Errorf("expected 1 book on page 2, got %d", len(booksPage2))
		}

		// Test search filtering
		booksSearch, _, _ := repo.FindAll(ctx, 1, 10, "Doe") // Should match "Book 3" by Doe
		if len(booksSearch) != 1 {
			t.Errorf("expected 1 book for search 'Doe', got %d", len(booksSearch))
		}
		if len(booksSearch) > 0 && booksSearch[0].Title != "Book 3" {
			t.Errorf("expected title 'Book 3', got '%s'", booksSearch[0].Title)
		}
	})

	// 3. Test Update
	t.Run("Update", func(t *testing.T) {
		book := domain.Book{ID: "1", Title: "Updated Go Book", Author: "John Doe", Year: 2023}
		err := repo.Update(ctx, book)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		found, _ := repo.FindByID(ctx, "1")
		if found.Title != "Updated Go Book" {
			t.Errorf("expected title 'Updated Go Book', got '%s'", found.Title)
		}
	})

	// 4. Test Delete
	t.Run("Delete", func(t *testing.T) {
		err := repo.Delete(ctx, "1")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		_, err = repo.FindByID(ctx, "1")
		if err == nil {
			t.Error("expected error 'book not found', got nil")
		}
	})
}
