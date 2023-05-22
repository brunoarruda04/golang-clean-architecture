package students

import (
	"net/http"

	"github.com/brunoarruda04/golang-clean-architecture/api/controller"
	"github.com/brunoarruda04/golang-clean-architecture/entities"
	"github.com/brunoarruda04/golang-clean-architecture/entities/shared"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var input Input
	var newStudents []entities.Student
	var err error

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.GetUuidByString(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problema com seu o ID"))
		return
	}

	for _, studentElement := range entities.Students {
		if input.UUID != studentElement.ID {
			newStudents = append(newStudents, studentElement)
		}
	}

	entities.Students = newStudents
	c.JSON(http.StatusOK, controller.NewReponseMessage("Estudante excluido com sucesso"))
}
