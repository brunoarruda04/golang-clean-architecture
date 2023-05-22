package infra

import (
	"net/http"

	"github.com/brunoarruda04/golang-clean-architecture/api/controller"
	"github.com/gin-gonic/gin"
)

func Heart(c *gin.Context) {
	c.JSON(http.StatusOK, controller.NewReponseMessage("OK"))
}
