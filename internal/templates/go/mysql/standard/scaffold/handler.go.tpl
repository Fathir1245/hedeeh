package handler

import (
	"encoding/json"
	"net/http"

	"{{.ModuleName}}/internal/service"
)

type UserHandler struct {
	userService UserService 
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	
}