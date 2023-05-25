package student

import (
	"errors"

	"github.com/brunoarruda04/golang-clean-architecture/entities/shared"
	"github.com/google/uuid"
)

func (su *StudentUsecase) Delete(id uuid.UUID) (err error) {
	student, err := su.Database.StudentRepository.FindById(id)

	if err != nil {
		return err
	}

	if student.ID == shared.GetUuidEmpty() {
		return errors.New("Não foi possivel encontrar o estudante")
	}

	return su.Database.StudentRepository.Delete(id)
}
