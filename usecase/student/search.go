package student

import (
	"errors"

	"github.com/brunoarruda04/golang-clean-architecture/entities"
	"github.com/brunoarruda04/golang-clean-architecture/entities/shared"
	"github.com/google/uuid"
)

func SearchById(id uuid.UUID) (student entities.Student, err error) {
	for _, studentElement := range entities.Students {
		if studentElement.ID == id {
			student = studentElement
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Estudante não encontrado")
	}

	return student, err
}
