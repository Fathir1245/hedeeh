package router

import (
	"net/http"
)

// RegisterUserRoutes registers all user-related routes with the provided ServeMux and handler.
func RegisterUserRoutes(mux *http.ServeMux, handler *UserHandler) {
	// POST   /user
	mux.HandleFunc("/user", handler.CreateUser)
	
	// GET    /user/{id}
	mux.HandleFunc("/user/{id}", handler.GetUser)
	
	// PUT    /user/{id}
	mux.HandleFunc("/user/{id}", handler.UpdateUser)
	
	// DELETE /user/{id}
	mux.HandleFunc("/user/{id}", handler.DeleteUser)
}