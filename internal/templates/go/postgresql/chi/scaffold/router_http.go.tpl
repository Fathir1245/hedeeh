package router

import (
	"net/http"

	"{{.ModuleName}}/internal/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	healthHandler := handler.NewHealthHandler()

	mux.HandleFunc("/health", healthHandler.Check)

	return mux
}
