package server

import (
	"fmt"

	"github/herochi/orbi/service-a/adapter/container"

	"github.com/gin-gonic/gin"
)

type ginServer struct {
	Router    *gin.Engine
	Container *container.Container
}

func NewServer(e *gin.Engine, c *container.Container) *ginServer {
	return &ginServer{
		Router:    e,
		Container: c,
	}
}

func (g *ginServer) Start(port string) error {
	return g.Router.Run(fmt.Sprintf(":%s", port))
}
