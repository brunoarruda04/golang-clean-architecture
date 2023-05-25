package students

import (
	"github.com/brunoarruda04/golang-clean-architecture/entities"
)

type StudentController struct {
	StudentsUsecase entities.StudentUsecaseContract
}

func NewStudentController(su entities.StudentUsecaseContract) *StudentController {
	return &StudentController{
		StudentsUsecase: su,
	}
}
