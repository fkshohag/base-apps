package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"xyz-task-2/internals/models"
)

// StudentService defines the interface for student-related operations
type StudentService interface {
	GetStudents() ([]models.Student, error)
	Create(student *models.Student) error
	Delete(id string) error
	GetByID(id string) (*models.Student, error)
	List(params *models.ListParams) ([]*models.Student, error)
	Update(id string, student *models.Student) error
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

func (h *StudentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *StudentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	student, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if student == nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) List(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters for filtering/pagination
	params := &models.ListParams{
		// You can add query parameter parsing here
		// For example:
		// Page: parseInt(r.URL.Query().Get("page")),
		// PageSize: parseInt(r.URL.Query().Get("pageSize")),
	}

	students, err := h.service.List(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func (h *StudentHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(id, &student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
