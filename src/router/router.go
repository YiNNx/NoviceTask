package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"src/controller"
	"src/utils"
)

func CreateRouters(e *echo.Echo) {
	g := e.Group("/user")
	g.POST("", controller.SignUP)
	g.GET("/token", controller.LogIn)

	g.GET("/:id", controller.GetUser, middleware.JWTWithConfig(utils.Conf), utils.VerifyUser)
	g.PUT("/:id", controller.ChangeInfo, middleware.JWTWithConfig(utils.Conf), utils.VerifyUser)

	g.GET("/all", controller.GetAllUser, middleware.JWTWithConfig(utils.Conf), utils.VerifyAdmin)
	g.DELETE("/:id", controller.DeleteUser, middleware.JWTWithConfig(utils.Conf), utils.VerifyAdmin)
}
