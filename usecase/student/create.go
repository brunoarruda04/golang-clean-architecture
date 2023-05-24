package student

import "github.com/brunoarruda04/golang-clean-architecture/entities"

func (su *StudentUsecase) Create(name string, age int) (entities.Student, error) {
	student := entities.NewStudent(name, age)
	err := su.Database.StudentRepository.Create(student)
	return *student, err
}
