package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"xyz-task-2/internals/models"
)

type StudentService interface {
	Create(student *models.Student) error
	GetByID(id string) (*models.Student, error)
	List(params *models.ListParams) ([]*models.Student, error)
	Update(id string, student *models.Student) error
	Delete(id string) error
}

type StudentHandler struct {
	service StudentService
}

func NewStudentHandler(service StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

func (h *StudentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // You might want to use a router that provides path params
	student, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) List(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	query := r.URL.Query()
	params := &models.ListParams{
		Page:     parseInt(query.Get("page"), 1),
		PageSize: parseInt(query.Get("pageSize"), 10),
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
	id := r.URL.Query().Get("id") // You might want to use a router that provides path params
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Update(id, &student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // You might want to use a router that provides path params
	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted successfully"})
}

// Helper function to parse integer parameters
func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return result
}
