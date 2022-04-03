package main

import (
	"github.com/labstack/echo/v4"
	"src/controllers"
	"src/models"
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
	//g := e.Group("/admin")
	//g.Use(middleware.BasicAuth(func(jwtToken string, _ string, c echo.Context) (bool, error) {
	//	if jwtToken == "foo" {
	//		return true, nil
	//	}
	//	return false, nil
	//}))
	e.Start(":8080")
}
