package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"xyz-task-2/internals/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of the users.Service
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(*MockUserService)
		expectedStatus int
		expectedBody   []models.User
	}{
		{
			name: "Success",
			setupMock: func(m *MockUserService) {
				m.On("GetUsers").Return([]models.User{
					{ID: "1", Username: "John Doe"},
					{ID: "2", Username: "Jane Doe"},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []models.User{
				{ID: "1", Username: "John Doe"},
				{ID: "2", Username: "Jane Doe"},
			},
		},
		{
			name: "Internal Server Error",
			setupMock: func(m *MockUserService) {
				m.On("GetUsers").Return([]models.User{}, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new mock service
			mockService := new(MockUserService)
			tt.setupMock(mockService)
			// Create handler with mock service
			handler := NewUserHandler(mockService)

			// Create test request
			req := httptest.NewRequest(http.MethodGet, "/users", nil)
			w := httptest.NewRecorder()

			// Call the handler
			handler.GetUsers(w, req)

			// Assert status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// If success case, verify response body
			if tt.expectedStatus == http.StatusOK {
				var response []models.User
				err := json.NewDecoder(w.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			}

			// Verify that all expected mock calls were made
			mockService.AssertExpectations(t)
		})
	}
}
