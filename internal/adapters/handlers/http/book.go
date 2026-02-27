package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"dsnt-challenge/internal/core/domain"
	"dsnt-challenge/internal/core/ports"
	"dsnt-challenge/pkg/logger"
	"dsnt-challenge/pkg/response"
)

type BooksHandler struct {
	service ports.BooksService
}

func NewBooksHandler(service ports.BooksService) *BooksHandler {
	return &BooksHandler{service: service}
}

func (h *BooksHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getBooks(w, r)
	case http.MethodPost:
		h.createBook(w, r)
	default:
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *BooksHandler) HandleBookByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from path, e.g., /books/id
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	if id == "" || id == r.URL.Path {
		response.Error(w, http.StatusBadRequest, "invalid book id")
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getBookByID(w, r, id)
	case http.MethodPut:
		h.updateBook(w, r, id)
	case http.MethodDelete:
		h.deleteBook(w, r, id)
	default:
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *BooksHandler) getBooks(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	searchStr := r.URL.Query().Get("search")

	page, _ := strconv.Atoi(pageStr) // Defaults format to 0 on err, handled by service
	limit, _ := strconv.Atoi(limitStr)

	req := domain.GetBooksRequest{
		Search: searchStr,
		Page:   page,
		Limit:  limit,
	}

	books, _, err := h.service.GetBooks(r.Context(), req)
	if err != nil {
		logger.Error("failed to get books", err)
		response.Error(w, http.StatusInternalServerError, "failed to get books")
		return
	}

	logger.Info("get books request handled successfully")
	response.JSON(w, http.StatusOK, books)
}

func (h *BooksHandler) getBookByID(w http.ResponseWriter, r *http.Request, id string) {
	book, err := h.service.GetBookByID(r.Context(), id)
	if err != nil {
		logger.Error("failed to get book by id", err, "id", id)
		if err.Error() == "book not found" {
			response.Error(w, http.StatusNotFound, err.Error())
		} else {
			response.Error(w, http.StatusInternalServerError, "failed to fetch book")
		}
		return
	}

	logger.Info("get book by id request handled successfully", "id", id)
	response.JSON(w, http.StatusOK, book)
}

func (h *BooksHandler) createBook(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("failed to decode create book request", err)
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	book, err := h.service.CreateBook(r.Context(), req)
	if err != nil {
		logger.Error("failed to create book", err, "req", req)
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	logger.Info("create book request handled successfully", "id", book.ID)
	response.JSON(w, http.StatusCreated, book)
}

func (h *BooksHandler) updateBook(w http.ResponseWriter, r *http.Request, id string) {
	var req domain.UpdateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("failed to decode update book request", err)
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	book, err := h.service.UpdateBook(r.Context(), id, req)
	if err != nil {
		logger.Error("failed to update book", err, "id", id)
		if err.Error() == "book not found" {
			response.Error(w, http.StatusNotFound, err.Error())
		} else {
			response.Error(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	logger.Info("update book request handled successfully", "id", book.ID)
	response.JSON(w, http.StatusOK, book)
}

func (h *BooksHandler) deleteBook(w http.ResponseWriter, r *http.Request, id string) {
	err := h.service.DeleteBook(r.Context(), id)
	if err != nil {
		logger.Error("failed to delete book", err, "id", id)
		if err.Error() == "book not found" {
			response.Error(w, http.StatusNotFound, err.Error())
		} else {
			response.Error(w, http.StatusInternalServerError, "failed to delete book")
		}
		return
	}

	logger.Info("delete book request handled successfully", "id", id)
	w.WriteHeader(http.StatusNoContent)
}
