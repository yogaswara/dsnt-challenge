package http

import (
	"dsnt-challenge/internal/core/domain"
	"dsnt-challenge/internal/core/ports"
	"dsnt-challenge/pkg/logger"
	"dsnt-challenge/pkg/response"
	"encoding/json"
	"net/http"
)

type EchoHandler struct {
	service ports.EchoService
}

func NewEchoHandler(service ports.EchoService) *EchoHandler {
	return &EchoHandler{service: service}
}

func (h *EchoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.EchoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("failed to decode request body", err)
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	res, err := h.service.Echo(r.Context(), req)
	if err != nil {
		logger.Error("echo service failed", err, "payload", req)
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	logger.Info("echo request handled successfully")
	response.JSON(w, http.StatusOK, res)
}
