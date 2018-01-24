package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"github.com/gotoolkit/cloudnativego/pkg/cloudnativego"
)

// Service represents a service for managing JWT tokens.
type Service struct {
	secret []byte
}

type claims struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

// NewService initializes a new service. It will generate a random key that will be used to sign JWT tokens.
func NewService() (*Service, error) {
	secret := securecookie.GenerateRandomKey(32)

	if secret == nil {
		return nil, cloudnativego.ErrSecretGeneration
	}
	service := &Service{
		secret,
	}
	return service, nil
}

// GenerateToken generates a new JWT token.
func (service *Service) GenerateToken(data *cloudnativego.TokenData) (string, error) {
	expireToken := time.Now().Add(time.Minute * 1).Unix()
	cl := claims{
		int(data.ID),
		data.Username,
		int(data.Role),
		jwt.StandardClaims{
			ExpiresAt: expireToken,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cl)
	signedToken, err := token.SignedString(service.secret)
	if err != nil {
		return "", err
	}
	return signedToken, err
}

// ParseAndVerifyToken parses a JWT token and verify its validity. It returns an error if token is invalid.
func (service *Service) ParseAndVerifyToken(token string) (*cloudnativego.TokenData, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return service.secret, nil
	})

	if err == nil && parsedToken != nil {
		if cl, ok := parsedToken.Claims.(*claims); ok && parsedToken.Valid {
			tokenData := &cloudnativego.TokenData{
				ID:       cloudnativego.UserID(cl.UserID),
				Username: cl.Username,
				Role:     cloudnativego.UserRole(cl.Role),
			}
			return tokenData, nil

		}
	}
	return nil, cloudnativego.ErrInvalidJWTToken
}
