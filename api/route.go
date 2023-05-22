package api

import (
	infra_controller "github.com/brunoarruda04/golang-clean-architecture/api/controller/infra"
	student_controlller "github.com/brunoarruda04/golang-clean-architecture/api/controller/students"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/heart", infra_controller.Heart)
	groupStudent := s.Engine.Group("students")
	groupStudent.GET("/", student_controlller.List)
	groupStudent.POST("/", student_controlller.Create)
	groupStudent.PUT("/:id", student_controlller.Update)
	groupStudent.DELETE("/:id", student_controlller.Delete)
	groupStudent.GET("/:id", student_controlller.Details)
}
