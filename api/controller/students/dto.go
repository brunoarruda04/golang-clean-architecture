package students

import "github.com/google/uuid"

type Input struct {
	ID   string
	UUID uuid.UUID
	Name string `json:"name"`
	Age  int    `json:"age"`
}
