package handlers

import (
	"testing"
	"xyz-task-2/internals/services/users"
)

func TestNewUserHandler(t *testing.T) {
	// Create a mock service
	mockService := &users.Service{}

	// Call the constructor
	handler := NewUserHandler(mockService)

	// Assert that handler is not nil
	if handler == nil {
		t.Error("Expected non-nil UserHandler")
	}

	// Assert that service is properly assigned
	if handler.service != mockService {
		t.Error("Service was not properly assigned to handler")
	}
}
