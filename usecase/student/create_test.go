package student

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func (s *StudentUsecaseSuite) TestCreateShouldReturnSuccess() {
	name := "teste"
	age := 18

	s.StudentRepository.EXPECT().Create(gomock.Any()).Return(nil).Times(1)

	student, err := s.StudentUsecase.Create(name, age)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), name, student.Name)
	assert.Equal(s.T(), age, student.Age)
}

func (s *StudentUsecaseSuite) TestCreateShouldReturnErr() {
	name := "teste"
	age := 18

	errExpect := fmt.Errorf("error expect")

	s.StudentRepository.EXPECT().Create(gomock.Any()).Return(errExpect).Times(1)

	student, err := s.StudentUsecase.Create(name, age)
	assert.Error(s.T(), err, errExpect.Error())
	assert.Equal(s.T(), name, student.Name)
	assert.Equal(s.T(), age, student.Age)
}
