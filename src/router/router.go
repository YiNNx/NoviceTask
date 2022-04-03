package router

import (
	"github.com/labstack/echo/v4"
	"src/controller"
)

func Router(e *echo.Echo) {
	e.POST("/user", controller.SignUP)
	e.GET("/user/token", controller.LogIn)
	e.GET("/user/:id", controller.GetUser)
	e.POST("/user/:id", controller.ChangeInfo)

	e.GET("/user/all", controller.GetAllUser)
	e.DELETE("/user/:id", controller.DeleteUser)
}
