package http

import (
	"dsnt-challenge/internal/core/ports"
	"dsnt-challenge/pkg/logger"
	"dsnt-challenge/pkg/response"
	"net/http"
)

type PingHandler struct {
	service ports.PingService
}

func NewPingHandler(service ports.PingService) *PingHandler {
	return &PingHandler{service: service}
}

func (h *PingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	logger.Info("ping request handled successfully")
	response.Success(w, http.StatusOK, "", nil)
}
