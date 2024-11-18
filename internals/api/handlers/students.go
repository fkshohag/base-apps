package handlers

import (
	"encoding/json"
	"net/http"

	"xyz-task-2/internals/models"
)

// StudentService defines the interface for student-related operations
type StudentService interface {
	GetStudents() ([]models.Student, error)
}

// StudentHandler struct now uses the interface instead of concrete type
type StudentHandler struct {
	service StudentService
}

// Update constructor to use interface
func NewStudentHandler(service StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.service.GetStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
