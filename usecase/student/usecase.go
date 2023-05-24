package student

import "github.com/brunoarruda04/golang-clean-architecture/infra/database"

type StudentUsecase struct {
	Database *database.Database
}

func NewStudentUsecase(db *database.Database) *StudentUsecase {
	return &StudentUsecase{
		Database: db,
	}
}
