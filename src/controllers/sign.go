package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"src/models"
)

type Response struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func SignUP(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return err
	}
	if err := u.Insert(); err != nil {
		return err
	}
	r := new(Response)
	r.Success = true
	r.Data = u
	return c.JSON(http.StatusOK, r)
}

func LogIn(c echo.Context) error {
	return nil
}

// route：e.GET("/user/:id", GetUser)
func GetUser(c echo.Context) error {
	//id, _ := strconv.Atoi(c.Param("id"))
	//u, err := models.SelectId(id)
	//if err != nil {
	//	return err
	//}
	//
	//return c.JSON(http.StatusOK, u)
	return nil
}

func ChangeInfo(c echo.Context) error {
	return nil
}

func GetAllUser(c echo.Context) error {
	return nil
}

func DeleteUser(c echo.Context) error {
	return nil
}
