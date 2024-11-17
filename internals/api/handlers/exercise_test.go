package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"xyz-task-2/internals/models"

	"github.com/stretchr/testify/assert"
)

// MockExerciseService implements ExerciseService interface
type MockExerciseService struct {
	exercise *models.ExerciseRecommendation
	err      error
}

func (m *MockExerciseService) GetExerciseRecommendation(userID string) (models.ExerciseRecommendation, error) {
	if m.exercise == nil {
		return models.ExerciseRecommendation{}, m.err
	}
	return *m.exercise, m.err
}

func TestGenerateExercise(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		mockService    *MockExerciseService
		expectedStatus int
		expectedBody   *models.ExerciseRecommendation
	}{
		{
			name:   "Success",
			userID: "123",
			mockService: &MockExerciseService{
				exercise: &models.ExerciseRecommendation{
					UserID:    "123",
					TopErrors: []models.Error{},
				},
				err: nil,
			},
			expectedStatus: http.StatusOK,
			expectedBody: &models.ExerciseRecommendation{
				UserID:    "123",
				TopErrors: []models.Error{},
			},
		},
		{
			name:           "Missing UserID",
			userID:         "",
			mockService:    &MockExerciseService{},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:   "Service Error",
			userID: "123",
			mockService: &MockExerciseService{
				exercise: nil,
				err:      errors.New("service error"),
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create handler with mock service
			handler := NewExerciseHandler(tt.mockService)

			// Create test request with query parameter
			req := httptest.NewRequest(http.MethodGet, "/exercise?user_id="+tt.userID, nil)
			w := httptest.NewRecorder()

			// Call the handler
			handler.GenerateExercise(w, req)

			// Assert status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// If success case, verify response body
			if tt.expectedStatus == http.StatusOK {
				var response models.ExerciseRecommendation
				err := json.NewDecoder(w.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, &response)
			}
		})
	}
}
