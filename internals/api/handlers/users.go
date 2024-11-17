package handlers

import (
	"encoding/json"
	"net/http"

	"xyz-task-2/internals/models"
)

// UserService defines the interface for user-related operations
type UserService interface {
	GetUsers() ([]models.User, error)
}

// UserHandler struct now uses the interface instead of concrete type
type UserHandler struct {
	service UserService
}

// Update constructor to use interface
func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
