package main

import (
	"github.com/labstack/echo/v4"
	"src/controller"
	"src/model"
)

func main() {
	model.Connect()
	defer model.Close()

	e := echo.New()
	e.POST("/user", controller.SignUP)
	e.GET("/user/token", controller.LogIn)
	e.GET("/user/:id", controller.GetUser)
	e.POST("/user/:id", controller.ChangeInfo)

	e.GET("/user/all", controller.GetAllUser)
	e.DELETE("/user/:id", controller.DeleteUser)

	// Restricted group
	//r := e.Group("/user")
	//// Configure middleware with the custom claims type
	//r.Use(middleware.JWTWithConfig(utils.Config))
	//r.GET("/t", utils.Restricted)

	e.Start(":8080")

}
