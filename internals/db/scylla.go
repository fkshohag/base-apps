package db

import (
	"fmt"
	"sort"
	"xyz-task-2/internals/models"

	"github.com/gocql/gocql"
)

type ScyllaClient struct {
	session *gocql.Session
}

type ScyllaConfig struct {
	Hosts    []string
	Keyspace string
}

func NewScyllaClient(cfg ScyllaConfig) (*ScyllaClient, error) {
	cluster := gocql.NewCluster(cfg.Hosts...)
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return &ScyllaClient{session: session}, nil
}

func (c *ScyllaClient) Execute(query string, values ...interface{}) error {
	return c.session.Query(query, values...).Exec()
}

func (c *ScyllaClient) Query(query string, values ...interface{}) *gocql.Iter {
	return c.session.Query(query, values...).Iter()
}

func (sc *ScyllaClient) GetSession() *gocql.Session {
	return sc.session
}

func (c *ScyllaClient) GetTopErrors(userID string, limit int) ([]models.Error, error) {
	var errors []models.Error
	query := `
		SELECT error_category, error_subcategory, frequency
		FROM error_frequencies
		WHERE user_id = ?
	`
	// Improvements needed, this is a bad approach
	iter := c.session.Query(query, userID).Iter()
	var category, subcategory string
	var frequency int
	for iter.Scan(&category, &subcategory, &frequency) {
		errors = append(errors, models.Error{
			Category:    category,
			Subcategory: subcategory,
			Frequency:   frequency,
		})
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	sort.Slice(errors, func(i, j int) bool {
		return errors[i].Frequency > errors[j].Frequency
	})

	if limit > len(errors) {
		limit = len(errors)
	}
	return errors[:limit], nil
}

func (c *ScyllaClient) GetUsers() ([]models.User, error) {
	var users []models.User
	query := `
		SELECT id, username FROM users;
	`
	iter := c.session.Query(query).Iter()
	var id string
	var username string
	fmt.Println("iteriteriteriteriteriteriteriteriteriteriteriteriteriteriteriter")
	for iter.Scan(&id, &username) {
		users = append(users, models.User{
			ID:       id,
			Username: username,
		})
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	fmt.Println(users)
	fmt.Println(":::::::::::::::::::::::::::::::::::::::::::::::::::")
	return users, nil
}

func (c *ScyllaClient) GetStudents() ([]models.Student, error) {
	var students []models.Student
	query := `
		SELECT id, name, department, roll, email, semester, batch_year FROM students;
	`
	iter := c.session.Query(query).Iter()
	var (
		id         string
		name       string
		department string
		roll       string
		email      string
		semester   int
		batchYear  int
	)

	for iter.Scan(&id, &name, &department, &roll, &email, &semester, &batchYear) {
		students = append(students, models.Student{
			ID:         id,
			Name:       name,
			Department: department,
			Roll:       roll,
			Email:      email,
			Semester:   semester,
			BatchYear:  batchYear,
		})
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return students, nil
}

func (c *ScyllaClient) Close() {
	c.session.Close()
}

func (c *ScyllaClient) CreateStudent(student models.Student) error {
	query := `
		INSERT INTO students (id, name, department, roll, email, semester, batch_year)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	return c.Execute(query,
		student.ID,
		student.Name,
		student.Department,
		student.Roll,
		student.Email,
		student.Semester,
		student.BatchYear,
	)
}

func (c *ScyllaClient) GetStudentByID(id string) (models.Student, error) {
	var student models.Student
	query := `
		SELECT id, name, department, roll, email, semester, batch_year
		FROM students
		WHERE id = ?
		LIMIT 1
	`
	iter := c.session.Query(query, id).Iter()
	var (
		name       string
		department string
		roll       string
		email      string
		semester   int
		batchYear  int
	)

	if !iter.Scan(&id, &name, &department, &roll, &email, &semester, &batchYear) {
		return student, fmt.Errorf("student not found with id: %s", id)
	}

	student = models.Student{
		ID:         id,
		Name:       name,
		Department: department,
		Roll:       roll,
		Email:      email,
		Semester:   semester,
		BatchYear:  batchYear,
	}

	if err := iter.Close(); err != nil {
		return student, err
	}

	return student, nil
}

func (c *ScyllaClient) UpdateStudent(student models.Student) error {
	query := `
		UPDATE students 
		SET name = ?, 
			department = ?, 
			roll = ?, 
			email = ?, 
			semester = ?, 
			batch_year = ?
		WHERE id = ?
	`
	return c.Execute(query,
		student.Name,
		student.Department,
		student.Roll,
		student.Email,
		student.Semester,
		student.BatchYear,
		student.ID,
	)
}

func (c *ScyllaClient) DeleteStudent(id string) error {
	query := `
		DELETE FROM students 
		WHERE id = ?
	`
	return c.Execute(query, id)
}
