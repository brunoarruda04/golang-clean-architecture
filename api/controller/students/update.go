package students

import (
	"net/http"

	"github.com/brunoarruda04/golang-clean-architecture/api/controller"
	"github.com/brunoarruda04/golang-clean-architecture/entities"
	"github.com/brunoarruda04/golang-clean-architecture/entities/shared"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var input Input
	var studentFound entities.Student
	var newStudents []entities.Student

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.GetUuidByString(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	for _, studentElement := range entities.Students {
		if studentElement.ID == input.UUID {
			studentFound = studentElement
		}
	}

	if studentFound.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Não foi possível encontrar o estudante"))
		return
	}

	studentFound.Name = input.Name
	studentFound.Age = input.Age

	for _, studentElement := range entities.Students {
		if studentFound.ID == studentElement.ID {
			newStudents = append(newStudents, studentFound)
		} else {
			newStudents = append(newStudents, studentElement)
		}
	}

	entities.Students = newStudents

	c.JSON(http.StatusOK, studentFound)
}
