package services

import (
	"context"
	"errors"
	"math"

	"dsnt-challenge/internal/core/domain"
	"dsnt-challenge/internal/core/ports"

	"github.com/google/uuid"
)

type booksService struct {
	repo ports.BookRepository
}

// NewBooksService creates a new instance of BooksService
func NewBooksService(repo ports.BookRepository) ports.BooksService {
	return &booksService{
		repo: repo,
	}
}

func (s *booksService) GetBooks(ctx context.Context, req domain.GetBooksRequest) ([]domain.Book, domain.PaginationMeta, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	books, total, err := s.repo.FindAll(ctx, req.Page, req.Limit, req.Search, req.Author)
	if err != nil {
		return nil, domain.PaginationMeta{}, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

	meta := domain.PaginationMeta{
		Page:       req.Page,
		Limit:      req.Limit,
		TotalItems: total,
		TotalPages: totalPages,
	}

	return books, meta, nil
}

func (s *booksService) GetBookByID(ctx context.Context, id string) (domain.Book, error) {
	if id == "" {
		return domain.Book{}, errors.New("id is required")
	}
	return s.repo.FindByID(ctx, id)
}

func (s *booksService) CreateBook(ctx context.Context, req domain.CreateBookRequest) (domain.Book, error) {
	if req.Title == "" || req.Author == "" {
		return domain.Book{}, errors.New("title and author are required")
	}

	newBook := domain.Book{
		ID:     uuid.New().String(),
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}

	if err := s.repo.Save(ctx, newBook); err != nil {
		return domain.Book{}, err
	}

	return newBook, nil
}

func (s *booksService) UpdateBook(ctx context.Context, id string, req domain.UpdateBookRequest) (domain.Book, error) {
	if id == "" {
		return domain.Book{}, errors.New("id is required")
	}
	if req.Title == "" || req.Author == "" {
		return domain.Book{}, errors.New("title and author are required")
	}

	existingBook, err := s.repo.FindByID(ctx, id)
	if err != nil {
		// If book doesn't exist, we Upsert it (create with provided ID)
		newBook := domain.Book{
			ID:     id,
			Title:  req.Title,
			Author: req.Author,
			Year:   req.Year,
		}
		if err := s.repo.Save(ctx, newBook); err != nil {
			return domain.Book{}, err
		}
		return newBook, nil
	}

	existingBook.Title = req.Title
	existingBook.Author = req.Author
	existingBook.Year = req.Year

	if err := s.repo.Update(ctx, existingBook); err != nil {
		return domain.Book{}, err
	}

	return existingBook, nil
}

func (s *booksService) DeleteBook(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	err := s.repo.Delete(ctx, id)
	if err != nil && err.Error() == "book not found" {
		// Idempotent delete: if already gone, we consider it success.
		return nil
	}
	return err
}
