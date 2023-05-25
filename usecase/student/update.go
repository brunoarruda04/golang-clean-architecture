package student

import (
	"errors"

	"github.com/brunoarruda04/golang-clean-architecture/entities"
	"github.com/brunoarruda04/golang-clean-architecture/entities/shared"
	"github.com/google/uuid"
)

func (su *StudentUsecase) Update(id uuid.UUID, name string, age int) (entities.Student, error) {
	student, err := su.Database.StudentRepository.FindById(id)
	if err != nil {
		return student, nil
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Não foi possivel encontrar o estudante")
	}

	student.Name = name
	student.Age = age

	err = su.Database.StudentRepository.Update(&student)
	return student, err
}
