package student

import (
	"errors"

	"github.com/brunoarruda04/golang-clean-architecture/entities"
	"github.com/brunoarruda04/golang-clean-architecture/entities/shared"
	"github.com/google/uuid"
)

func UpdateById(id uuid.UUID, name string, age int) (student entities.Student, err error) {
	var newStudents []entities.Student

	for _, studentElement := range entities.Students {
		if studentElement.ID == id {
			student = studentElement
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Não foi possivel encontrar o estudante")
	}

	student.Name = name
	student.Age = age

	for _, studentElement := range entities.Students {
		if student.ID == studentElement.ID {
			newStudents = append(newStudents, student)
		} else {
			newStudents = append(newStudents, studentElement)
		}
	}

	entities.Students = newStudents
	return student, err
}
