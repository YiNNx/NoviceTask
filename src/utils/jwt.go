package utils

import (
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtUserClaims struct {
	Id   int  `json:"id"`
	Role bool `json:"role"`
	jwt.StandardClaims
}

var Config = middleware.JWTConfig{
	Claims:     &jwtUserClaims{},
	SigningKey: []byte("secret"),
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
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ""
	}
	return t
}

func Restricted(c echo.Context) error {
	user := c.Get("id").(*jwt.Token)
	claims := user.Claims.(*jwtUserClaims)
	name := claims.Id
	return c.String(http.StatusOK, "Welcome "+string(name)+"!")
}
