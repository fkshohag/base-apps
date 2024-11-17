package handlers

import (
	"encoding/json"
	"net/http"

	"xyz-task-2/internals/models"
)

// ExerciseService defines the interface for exercise-related operations
type ExerciseService interface {
	GetExerciseRecommendation(userID string) (models.ExerciseRecommendation, error)
}

// ExerciseHandler struct
type ExerciseHandler struct {
	service ExerciseService
}

// NewExerciseHandler constructor
func NewExerciseHandler(service ExerciseService) *ExerciseHandler {
	return &ExerciseHandler{service: service}
}

func (h *ExerciseHandler) GenerateExercise(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	recommendation, err := h.service.GetExerciseRecommendation(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendation)
}
