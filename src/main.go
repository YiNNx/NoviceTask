package main

import (
	"github.com/labstack/echo/v4"
	"src/controllers"
	"src/models"
)

func main() {
	e := echo.New()
	models.Connect()
	defer models.Close()
	//models.CreateSchema()
	e.POST("/signup", controllers.SignUP)
	e.GET("/user/:id", controllers.GetUser)
	e.Start(":8080")
}
