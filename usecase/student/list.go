package student

import "github.com/brunoarruda04/golang-clean-architecture/entities"

func (su *StudentUsecase) List() (students []entities.Student, err error) {
	students = entities.Students
	return su.Database.StudentRepository.List()
}
