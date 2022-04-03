package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"src/models"
	"strconv"
)

type Response struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type passport struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

func ResJSON200(c echo.Context, msg string, data interface{}) error {
	return c.JSON(
		http.StatusOK,
		Response{true, msg, data})
}

func ResJSON400(c echo.Context, msg string) error {
	return c.JSON(
		http.StatusBadRequest,
		Response{false, msg, nil})
}

func ResJSON401(c echo.Context, msg string) error {
	return c.JSON(
		http.StatusUnauthorized,
		Response{false, msg, nil})
}

func ResJSON500(c echo.Context, msg string) error {
	return c.JSON(
		http.StatusInternalServerError,
		Response{false, msg, nil})
}

func SignUP(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return ResJSON400(c, err.Error())
	}
	if err := u.Insert(); err != nil {
		return ResJSON401(c, err.Error())
	}
	return ResJSON200(c, "registered successfully", u)
}

func LogIn(c echo.Context) error {
	p := new(passport)

	if err := c.Bind(p); err != nil {
		return ResJSON400(c, err.Error())
	}

	res, err := models.CheckUser(p.Email, p.Pwd)
	if err != nil {
		return ResJSON400(c, err.Error())
	}
	if !res {
		return ResJSON401(c, err.Error())
	}
	return ResJSON200(c, "logged in successfully", p)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := models.GetUser(id)
	if err != nil {
		return ResJSON400(c, err.Error())
	}
	return ResJSON200(c, "data query succeeded", u)
}

func ChangeInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return ResJSON400(c, err.Error())
	}
	user, err := u.Update(id)
	if err != nil {
		return ResJSON401(c, err.Error())
	}
	return ResJSON200(c, "data updated successfully", user)
}

func GetAllUser(c echo.Context) error {
	users, err := models.SelectAllUser()
	if err != nil {
		return ResJSON500(c, err.Error())
	}
	return ResJSON200(c, "data query succeeded", users)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteUser(id)
	if err != nil {
		return ResJSON401(c, err.Error())
	}
	return ResJSON200(c, "deleted successfully", nil)
}
