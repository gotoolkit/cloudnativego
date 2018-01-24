package handler

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/cloudnativego/pkg/cloudnativego"
)

// AuthHandler represents an HTTP API handler for managing authentication.
type AuthHandler struct {
	UserService cloudnativego.UserService
	JWTService  cloudnativego.JWTService
}

// NewAuth returns a new instance of AuthHandler.
func NewAuth(rg *gin.RouterGroup) *AuthHandler {
	h := &AuthHandler{}
	rg.POST("/", h.auth)
	return h
}

type (
	postAuthRequest struct {
		Username string `valid:"required"`
		Password string `valid:"required"`
	}
)

func (handler *AuthHandler) auth(c *gin.Context) {
	var req postAuthRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	username := req.Username

	u, err := handler.UserService.UserByUsername(username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	tokenData := &cloudnativego.TokenData{
		ID:       u.ID,
		Username: u.Username,
		Role:     u.Role,
	}
	token, err := handler.JWTService.GenerateToken(tokenData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
