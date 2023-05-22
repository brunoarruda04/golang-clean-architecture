package students

import (
	"net/http"

	student_usecase "github.com/brunoarruda04/golang-clean-architecture/usecase/student"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	students, err := student_usecase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, students)
}
