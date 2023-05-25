package students

import (
	"net/http"

	"github.com/brunoarruda04/golang-clean-architecture/api/controller"
	"github.com/brunoarruda04/golang-clean-architecture/entities/shared"
	"github.com/gin-gonic/gin"
)

func (sc *StudentController) Delete(c *gin.Context) {
	var input Input
	var err error

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.GetUuidByString(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problema com seu o ID"))
		return
	}
	err = sc.StudentsUsecase.Delete(input.UUID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, controller.NewReponseMessage("Estudante excluido com sucesso"))
}
