package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"src/models"
)

func SignUP(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return err
	}
	if err := u.InsertUser(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
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
