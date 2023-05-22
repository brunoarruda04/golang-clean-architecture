package student

import "github.com/brunoarruda04/golang-clean-architecture/entities"

func Create(name string, age int) (student entities.Student, err error) {
	student = *(entities.NewStudent(name, age))
	entities.Students = append(entities.Students, student)
	return student, err
}
