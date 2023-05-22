package students

import (
	"net/http"

	"github.com/brunoarruda04/golang-clean-architecture/api/controller"
	"github.com/brunoarruda04/golang-clean-architecture/entities"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var input Input
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student := entities.NewStudent(input.Name, input.Age)
	entities.Students = append(entities.Students, *student)

	c.JSON(http.StatusAccepted, student)
}
