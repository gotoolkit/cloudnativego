package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/cloudnativego/pkg/cloudnativego"
)

// Middleware represents an entity that manages all middleware
type Middleware struct {
	jwtService cloudnativego.JWTService
}

// New initializes middleware
func New(jwtService cloudnativego.JWTService) *Middleware {
	return &Middleware{
		jwtService: jwtService,
	}
}

// JWTAuth provides Authentication middleware for handlers
func (mw *Middleware) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenData *cloudnativego.TokenData
		var token string
		tokens, ok := c.Request.Header["Authorization"]
		if ok && len(tokens) >= 1 {
			token = tokens[0]
			token = strings.TrimPrefix(token, "Bearer ")
		}
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": cloudnativego.ErrUnauthorized})
			return
		}
		var err error
		tokenData, err = mw.jwtService.ParseAndVerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
			return
		}
		storeTokenData(c, tokenData)
		c.Next()
	}
}

func storeTokenData(c *gin.Context, tokenData *cloudnativego.TokenData) {

}
