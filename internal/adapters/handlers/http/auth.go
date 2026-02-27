package http

import (
	"encoding/json"
	"net/http"

	"dsnt-challenge/internal/core/domain"
	"dsnt-challenge/internal/core/ports"
	"dsnt-challenge/pkg/logger"
	"dsnt-challenge/pkg/response"
)

type AuthHandler struct {
	service ports.AuthService
}

func NewAuthHandler(service ports.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) HandleToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("failed to decode auth request", err)
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	res, err := h.service.Login(r.Context(), req)
	if err != nil {
		logger.Error("login failed", err)
		response.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	logger.Info("login request handled successfully")
	// Return as raw JSON without standard response wrapper so tester can easily read token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
