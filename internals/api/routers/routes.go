package routers

import (
	"net/http"
	"xyz-task-2/internals/api/handlers"
	"xyz-task-2/internals/db"
	"xyz-task-2/internals/middlewares"
	"xyz-task-2/internals/services/recommendation"
	"xyz-task-2/internals/services/students"
	"xyz-task-2/internals/services/users"

	"github.com/gorilla/mux"
)

type responseWriterWrapper struct {
	http.ResponseWriter
}

func SetupRoutes(scyllaClient *db.ScyllaClient, redisClient *db.RedisClient) *mux.Router {
	router := mux.NewRouter()

	recommendationService := recommendation.NewService(scyllaClient, redisClient)
	userService := users.NewService(scyllaClient, redisClient)
	studentService := students.NewService(scyllaClient)

	exerciseHandler := handlers.NewExerciseHandler(recommendationService)
	usersHandler := handlers.NewUserHandler(userService)
	studentsHandler := handlers.NewStudentHandler(studentService)
	healthHandler := handlers.NewHealthHandler()

	router.Use(middlewares.Logging)
	router.Use(middlewares.CORS)

	router.HandleFunc("/health", healthHandler.Check).Methods("GET")

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/generate-exercise", exerciseHandler.GenerateExercise).Methods("GET")
	api.HandleFunc("/users", usersHandler.GetUsers).Methods("GET")
	api.HandleFunc("/students", studentsHandler.List).Methods("GET")
	api.HandleFunc("/students/{id}", studentsHandler.GetByID).Methods("GET")
	api.HandleFunc("/students", studentsHandler.Create).Methods("POST")
	api.HandleFunc("/students/{id}", studentsHandler.Update).Methods("PUT")
	api.HandleFunc("/students/{id}", studentsHandler.Delete).Methods("DELETE")

	return router
}
