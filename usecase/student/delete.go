package student

import (
	"github.com/brunoarruda04/golang-clean-architecture/entities"
	"github.com/google/uuid"
)

func DeleteById(id uuid.UUID) (err error) {
	var newStudents []entities.Student
	for _, studentElement := range entities.Students {
		if id != studentElement.ID {
			newStudents = append(newStudents, studentElement)
		}
	}

	entities.Students = newStudents
	return err
}
