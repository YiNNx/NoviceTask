package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"src/controllers"
	"src/models"
	"src/utils"
)

func main() {
	models.Connect()
	defer models.Close()

	e := echo.New()
	e.Debug = true
	e.POST("/signup", controllers.SignUP)
	e.POST("/login", controllers.LogIn)
	e.GET("/user/:id", controllers.GetUser)
	e.POST("/user/:id", controllers.ChangeInfo)

	e.GET("/user/all", controllers.GetAllUser)
	e.DELETE("/user/:id", controllers.DeleteUser)

	// Restricted group
	r := e.Group("/test")

	// Configure middleware with the custom claims type
	r.Use(middleware.JWTWithConfig(utils.Config))
	r.GET("/t", utils.Restricted)
	e.Start(":8080")
}
