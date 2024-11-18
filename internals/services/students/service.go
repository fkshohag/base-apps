package students

import (
	"xyz-task-2/internals/db"
	"xyz-task-2/internals/models"
)

// Service struct now uses the interface instead of concrete implementation
type Service struct {
	scyllaClient *db.ScyllaClient
	redisClient  *db.RedisClient
}

// Create implements handlers.StudentService.
func (s *Service) Create(student *models.Student) error {
	err := s.scyllaClient.CreateStudent(*student)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements handlers.StudentService.
func (s *Service) Delete(id string) error {
	err := s.scyllaClient.DeleteStudent(id)
	if err != nil {
		return err
	}

	// Invalidate both specific student and list caches
	return nil
}

// GetByID implements handlers.StudentService.
func (s *Service) GetByID(id string) (*models.Student, error) {
	// If not in cache, get from DB
	student, err := s.scyllaClient.GetStudentByID(id)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

// List implements handlers.StudentService.
func (s *Service) List(params *models.ListParams) ([]*models.Student, error) {
	// If not in cache, get from DB
	students, err := s.scyllaClient.GetStudents()
	if err != nil {
		return nil, err
	}

	// Convert []models.Student to []*models.Student
	studentPtrs := make([]*models.Student, len(students))
	for i := range students {
		studentPtrs[i] = &students[i]
	}

	return studentPtrs, nil
}

// Update implements handlers.StudentService.
func (s *Service) Update(id string, student *models.Student) error {
	// Ensure the ID matches
	student.ID = id

	err := s.scyllaClient.UpdateStudent(*student)
	if err != nil {
		return err
	}

	return nil
}

// NewService now accepts the interface
func NewService(scyllaClient *db.ScyllaClient, redisClient *db.RedisClient) *Service {
	return &Service{
		scyllaClient: scyllaClient,
		redisClient:  redisClient,
	}
}

func (s *Service) GetStudents() ([]models.Student, error) {
	students, err := s.scyllaClient.GetStudents()
	if err != nil {
		return nil, err
	}
	return students, nil
}
