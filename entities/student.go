package entities

import (
	"github.com/brunoarruda04/golang-clean-architecture/entities/shared"
	"github.com/google/uuid"
)

type Student struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

func NewStudent(name string, age int) *Student {
	return &Student{
		ID:   shared.GetUuid(),
		Name: name,
		Age:  age,
	}
}

var Students = []Student{
	Student{ID: shared.GetUuid(), Name: "Bruno", Age: 18},
	Student{ID: shared.GetUuid(), Name: "Gabriela", Age: 18},
}
