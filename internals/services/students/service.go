package students

import (
	"xyz-task-2/internals/models"
)

// StudentRepository defines the interface for student data operations
type StudentRepository interface {
	GetStudents() ([]models.Student, error)
}

// Service struct now uses the interface instead of concrete implementation
type Service struct {
	repository StudentRepository
}

// Create implements handlers.StudentService.
func (s *Service) Create(student *models.Student) error {
	panic("unimplemented")
}

// Delete implements handlers.StudentService.
func (s *Service) Delete(id string) error {
	panic("unimplemented")
}

// GetByID implements handlers.StudentService.
func (s *Service) GetByID(id string) (*models.Student, error) {
	panic("unimplemented")
}

// List implements handlers.StudentService.
func (s *Service) List(params *models.ListParams) ([]*models.Student, error) {
	panic("unimplemented")
}

// Update implements handlers.StudentService.
func (s *Service) Update(id string, student *models.Student) error {
	panic("unimplemented")
}

// NewService now accepts the interface
func NewService(repository StudentRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetStudents() ([]models.Student, error) {
	return s.repository.GetStudents()
}
