package students

import (
	"net/http"

	"github.com/brunoarruda04/golang-clean-architecture/api/controller"
	"github.com/brunoarruda04/golang-clean-architecture/entities/shared"
	student_usecase "github.com/brunoarruda04/golang-clean-architecture/usecase/student"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var input Input
	var err error

	err = c.Bind(&input)
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

	student, err := student_usecase.UpdateById(input.UUID, input.Name, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, student)
}
