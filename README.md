# DSNT Challenge Context

This project is a RESTful API built with Go, implementing a strict Hexagonal Architecture (Ports and Adapters). It provides standard responses, error logging, and high test coverage.

## Tech Stack
- **Go**: >= `1.21` (Uses standard library for routing `net/http`)
- **Logging**: `log/slog` (Standard structured logger in Go 1.21+)
- **Testing**: `testing` and `net/http/httptest` (Standard library)
- **Architecture**: Hexagonal Architecture

## Architecture Structure
```text
project_root/
├── cmd/
│   └── api/
│       └── main.go       (Entry point)
├── internal/
│   ├── core/
│   │   ├── domain/       (Business entities, models)
│   │   ├── ports/        (Interfaces)
│   │   └── services/     (Business logic implementation - use cases)
│   └── adapters/
│       └── handlers/     (HTTP handlers/controllers - driving adapters)
│           └── http/
└── pkg/
    ├── response/         (Standardized response formats)
    └── logger/           (Error logging)
```

## Features
- **Hexagonal Architecture**: Clear separation of concerns (Domain, Ports, Services, Handlers).
- **Unit Tests**: Full coverage for business logic and HTTP handlers.
- **Error Logging**: Uses core Go structured logging.
- **Standardized Responses**: Consistent JSON payload structures for success and error.

## Endpoints

### 1. Ping
- **GET** `/ping`
- **Response**: 
```json
{
  "success": true
}
```

### 2. Echo
- **POST** `/echo`
- **Request Body**:
```json
{
  "message": "hello world"
}
```
- **Response**: 
```json
{
  "message": "hello world"
}
```
- **Error Response**:
```json
{
  "success": false,
  "error": "message cannot be empty"
}
```

### 3. Books CRUD
- **POST** `/books`: Create a new book. Body: `{"title": "Book Title", "author": "Author Name", "year": 2024}`.
- **GET** `/books`: List all books with pagination and search. Query parameters: `?page=1&limit=10&search=keyword`.
  Response includes pagination metadata:
  ```json
  {
    "success": true,
    "message": "Successfully fetched books",
    "data": [...],
    "meta": {
      "page": 1,
      "limit": 10,
      "total_items": 1,
      "total_pages": 1
    }
  }
  ```
- **GET** `/books/:id`: Get a book by its `id`.
- **PUT** `/books/:id`: Update an existing book's details. Body: `{"title": "New", "author": "New", "year": 2025}`.
- **DELETE** `/books/:id`: Delete a book.

## How to Run

1. **Clone & Navigate**
   ```bash
   cd dsnt-challenge
   ```

2. **Download Dependencies** (It strictly uses Go Standard Libraries)
   ```bash
   go mod tidy
   ```

3. **Run the API**
   ```bash
   go run cmd/api/main.go
   ```
   The server will start on port `8080`.

## How to Test

To run all unit tests for the application:
```bash
go test -v ./...
```
