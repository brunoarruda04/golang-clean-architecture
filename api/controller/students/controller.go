package students

import usecase "github.com/brunoarruda04/golang-clean-architecture/usecase/student"

type StudentController struct {
	StudentsUsecase *usecase.StudentUsecase
}

func NewStudentController(su *usecase.StudentUsecase) *StudentController {
	return &StudentController{
		StudentsUsecase: su,
	}
}
