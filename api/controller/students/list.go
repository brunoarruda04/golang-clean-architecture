package students

import (
	"net/http"

	"github.com/brunoarruda04/golang-clean-architecture/entities"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, entities.Students)
}
