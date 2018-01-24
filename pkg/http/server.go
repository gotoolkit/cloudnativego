package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/cloudnativego/pkg/cloudnativego"
	"github.com/gotoolkit/cloudnativego/pkg/http/handler"
	"github.com/gotoolkit/cloudnativego/pkg/http/middleware"
)

// Server implements the cloudnativego.Server interface
type Server struct {
	BindAddress string
	UserService cloudnativego.UserService
	JWTService  cloudnativego.JWTService
}

// Start starts the HTTP server
func (server *Server) Start() error {
	mw := middleware.New(server.JWTService)

	engine := gin.Default()
	v1 := engine.Group("/api/v1")
	authHandler := handler.NewAuth(v1.Group("/auth"))
	authHandler.UserService = server.UserService
	authHandler.JWTService = server.JWTService
	userHandler := handler.NewUser(v1.Group("/users", mw.JWTAuth()))
	userHandler.UserService = server.UserService

	return engine.Run(server.BindAddress)
}
