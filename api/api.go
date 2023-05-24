package api

import (
	"fmt"

	"github.com/brunoarruda04/golang-clean-architecture/api/controller/students"
	"github.com/brunoarruda04/golang-clean-architecture/infra/config"
	"github.com/brunoarruda04/golang-clean-architecture/infra/database"
	"github.com/brunoarruda04/golang-clean-architecture/usecase/student"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Engine            *gin.Engine
	Database          *database.Database
	StudentController *students.StudentController
}

func NewService(db *database.Database) *Service {
	return &Service{
		Engine:   gin.Default(),
		Database: db,
	}
}

func (s *Service) GetControllers() {
	studentUsecase := student.NewStudentUsecase(s.Database)
	s.StudentController = students.NewStudentController(studentUsecase)
}

func (s *Service) Start() error {
	s.GetControllers()
	s.GetRoutes()
	return s.Engine.Run(fmt.Sprintf(":%d", config.Env.ApiPort))
}
