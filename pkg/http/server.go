package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/cloudnativego/pkg/cloudnativego"
	"github.com/gotoolkit/cloudnativego/pkg/http/handler"
)

type Server struct {
	BindAddress string
	UserService cloudnativego.UserService
}

func (server *Server) Start() error {

	engine := gin.Default()

	userHandler := handler.NewUserHandler(engine.Group("/api/v1/todos"))
	userHandler.UserService = server.UserService

	return engine.Run(server.BindAddress)
}
