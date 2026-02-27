package ports

import (
	"context"
	"dsnt-challenge/internal/core/domain"
)

// PingService defines the interface for ping related operations
type PingService interface {
	Ping(ctx context.Context) string
}

// EchoService defines the interface for echo related operations
type EchoService interface {
	Echo(ctx context.Context, req domain.EchoRequest) (domain.EchoResponse, error)
}

// BookRepository defines the interface for database operations related to books
type BookRepository interface {
	FindAll(ctx context.Context, page, limit int, search string) ([]domain.Book, int, error)
	FindByID(ctx context.Context, id string) (domain.Book, error)
	Save(ctx context.Context, book domain.Book) error
	Update(ctx context.Context, book domain.Book) error
	Delete(ctx context.Context, id string) error
}

// BooksService defines the interface for books related operations
type BooksService interface {
	GetBooks(ctx context.Context, req domain.GetBooksRequest) ([]domain.Book, domain.PaginationMeta, error)
	GetBookByID(ctx context.Context, id string) (domain.Book, error)
	CreateBook(ctx context.Context, req domain.CreateBookRequest) (domain.Book, error)
	UpdateBook(ctx context.Context, id string, req domain.UpdateBookRequest) (domain.Book, error)
	DeleteBook(ctx context.Context, id string) error
}
