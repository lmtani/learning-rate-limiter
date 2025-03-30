package web

import (
	"encoding/json"
	"net/http"

	"github.com/lmtani/learning-rate-limiter/internal/usecase"
)

type HelloHandler struct {
	HelloUseCase *usecase.HelloUseCase
}

func NewHelloHandler(helloUseCase *usecase.HelloUseCase) *HelloHandler {
	return &HelloHandler{
		HelloUseCase: helloUseCase,
	}
}

func (h *HelloHandler) Handle(w http.ResponseWriter, r *http.Request) {
	output := h.HelloUseCase.Execute()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
