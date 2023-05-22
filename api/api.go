package api

import (
	"fmt"

	"github.com/brunoarruda04/golang-clean-architecture/infra/config"
	"github.com/gin-gonic/gin"
)

type Service struct {
	*gin.Engine
}

func NewService() *Service {
	return &Service{
		gin.Default(),
	}
}

func (s *Service) Start() error {
	s.GetRoutes()
	return s.Engine.Run(fmt.Sprintf(":%d", config.Env.ApiPort))
}
