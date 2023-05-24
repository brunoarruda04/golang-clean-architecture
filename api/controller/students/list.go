package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (sc *StudentController) List(c *gin.Context) {
	students, err := sc.StudentsUsecase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, students)
}
