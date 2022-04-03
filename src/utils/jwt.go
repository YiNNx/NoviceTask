package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"src/config"
	"time"
)

type jwtUserClaims struct {
	Id   int  `json:"id"`
	Role bool `json:"role"`
	jwt.StandardClaims
}

var Conf = middleware.JWTConfig{
	Claims:     &jwtUserClaims{},
	SigningKey: []byte(config.JwtSecret),
}

func GenerateToken(id int, role bool) string {
	claims := &jwtUserClaims{
		id,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "error"
	}

	return t
}
