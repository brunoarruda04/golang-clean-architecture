package students

import (
	"net/http"

	"github.com/brunoarruda04/golang-clean-architecture/api/controller"
	student_usecase "github.com/brunoarruda04/golang-clean-architecture/usecase/student"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var input Input
	var err error
	if err = c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := student_usecase.Create(input.Name, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}
	c.JSON(http.StatusAccepted, student)
}
