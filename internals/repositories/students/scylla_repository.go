package students

import (
	"xyz-task-2/internals/db"
	"xyz-task-2/internals/models"
	"xyz-task-2/internals/services/students"
)

type ScyllaRepository struct {
	client *db.ScyllaClient
}

func NewScyllaRepository(client *db.ScyllaClient) students.StudentRepository {
	return &ScyllaRepository{
		client: client,
	}
}

func (r *ScyllaRepository) GetStudents() ([]models.Student, error) {
	return r.client.GetStudents()
}
