package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"src/models"
)

func SignUP(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusUnauthorized, nil)
	}
	//user1 := &models.User{
	//	Name:  "admin",
	//	Email: "admin1@admin",
	//}
	models.InsertUser(u)
	return c.JSON(http.StatusOK, u)
}

// route：e.GET("/users/:id", GetUser)
func GetUser(c echo.Context) error {
	// 获取url上的path参数，url模式里面定义了参数:id
	id := c.Param("id")
	//响应一个字符串，这里直接把id以字符串的形式返回给客户端。
	username := c.QueryParam("username") //值为："tizi365"

	return c.String(http.StatusOK, id+" "+username)
}
