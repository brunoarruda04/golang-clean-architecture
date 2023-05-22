package students

import (
	"net/http"

	"github.com/brunoarruda04/golang-clean-architecture/api/controller"
	"github.com/brunoarruda04/golang-clean-architecture/entities"
	"github.com/brunoarruda04/golang-clean-architecture/entities/shared"
	"github.com/gin-gonic/gin"
)

func Details(c *gin.Context) {
	var student entities.Student
	var input Input
	var err error

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.GetUuidByString(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problema com seu o ID"))
		return
	}

	for _, studentElement := range entities.Students {
		if input.UUID == studentElement.ID {
			student = studentElement
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Estudante não encontrado"))
		return
	}

	c.JSON(http.StatusOK, student)
}
