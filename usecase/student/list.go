package student

import "github.com/brunoarruda04/golang-clean-architecture/entities"

func List() (students []entities.Student, err error) {
	students = entities.Students
	return students, err
}
