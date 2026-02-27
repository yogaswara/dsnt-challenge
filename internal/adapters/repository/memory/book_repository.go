package memory

import (
	"context"
	"errors"
	"sort"
	"strings"
	"sync"

	"dsnt-challenge/internal/core/domain"
	"dsnt-challenge/internal/core/ports"
)

type bookRepository struct {
	sync.RWMutex
	books map[string]domain.Book
}

// NewBookRepository creates a new instance of in-memory BookRepository
func NewBookRepository() ports.BookRepository {
	return &bookRepository{
		books: make(map[string]domain.Book),
	}
}

func (r *bookRepository) FindAll(ctx context.Context, page, limit int, search string) ([]domain.Book, int, error) {
	r.RLock()
	defer r.RUnlock()

	var filteredBooks []domain.Book
	searchLower := strings.ToLower(search)

	for _, b := range r.books {
		// Filter by search string in title or author (case-insensitive)
		if search == "" ||
			strings.Contains(strings.ToLower(b.Title), searchLower) ||
			strings.Contains(strings.ToLower(b.Author), searchLower) {
			filteredBooks = append(filteredBooks, b)
		}
	}

	total := len(filteredBooks)

	// Default to returning all if page/limit is non-positive or very large
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	// Sort map keys for deterministic pagination (by ID)
	sort.Slice(filteredBooks, func(i, j int) bool {
		return filteredBooks[i].ID < filteredBooks[j].ID
	})

	start := (page - 1) * limit
	if start >= total {
		return []domain.Book{}, total, nil
	}

	end := start + limit
	if end > total {
		end = total
	}

	return filteredBooks[start:end], total, nil
}

func (r *bookRepository) FindByID(ctx context.Context, id string) (domain.Book, error) {
	r.RLock()
	defer r.RUnlock()

	book, exists := r.books[id]
	if !exists {
		return domain.Book{}, errors.New("book not found")
	}

	return book, nil
}

func (r *bookRepository) Save(ctx context.Context, book domain.Book) error {
	r.Lock()
	defer r.Unlock()

	r.books[book.ID] = book
	return nil
}

func (r *bookRepository) Update(ctx context.Context, book domain.Book) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.books[book.ID]; !exists {
		return errors.New("book not found")
	}

	r.books[book.ID] = book
	return nil
}

func (r *bookRepository) Delete(ctx context.Context, id string) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.books[id]; !exists {
		return errors.New("book not found")
	}

	delete(r.books, id)
	return nil
}
