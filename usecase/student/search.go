package student

import (
	"github.com/brunoarruda04/golang-clean-architecture/entities"
	"github.com/google/uuid"
)

func (su *StudentUsecase) SearchById(id uuid.UUID) (entities.Student, error) {
	return su.Database.StudentRepository.FindById(id)
}
