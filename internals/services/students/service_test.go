package students

import (
	"testing"
	"xyz-task-2/internals/models"
)

// MockStudentRepository implements StudentRepository for testing
type MockStudentRepository struct {
	students []models.Student
	err      error
}

func (m *MockStudentRepository) GetStudents() ([]models.Student, error) {
	return m.students, m.err
}

func TestGetStudents(t *testing.T) {
	// Arrange
	mockStudents := []models.Student{
		{
			ID:         "1",
			Name:       "John Doe",
			Department: "CS",
			Roll:       "CS001",
			Email:      "john@example.com",
			Semester:   1,
			BatchYear:  2023,
		},
	}

	mockRepo := &MockStudentRepository{
		students: mockStudents,
		err:      nil,
	}

	service := NewService(mockRepo)

	// Act
	students, err := service.GetStudents()

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(students) != len(mockStudents) {
		t.Errorf("Expected %d students, got %d", len(mockStudents), len(students))
	}

	if students[0].ID != mockStudents[0].ID {
		t.Errorf("Expected student ID %s, got %s", mockStudents[0].ID, students[0].ID)
	}
}
