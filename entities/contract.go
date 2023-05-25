package entities

import "github.com/google/uuid"

type StudentUsecaseContract interface {
	Create(name string, age int) (Student, error)
	Delete(id uuid.UUID) error
	List() ([]Student, error)
	SearchById(id uuid.UUID) (Student, error)
	Update(id uuid.UUID, name string, age int) (Student, error)
}
