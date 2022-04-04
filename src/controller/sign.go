package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"src/model"
	"src/utils"
	"strconv"
	"time"
)

type receiveSignUp struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,max=20,min=4"`
	Pwd      string `json:"pwd" validate:"required,max=20,min=6"`
}

type responseToken struct {
	Token string `json:"token"`
}

func SignUP(c echo.Context) error {
	receive := new(receiveSignUp)
	if err := c.Bind(receive); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(receive); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	pwdHash, err := utils.PwdHash(receive.Pwd)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	u := &model.User{
		Email:    receive.Email,
		Username: receive.Username,
		PwdHash:  string(pwdHash),
	}
	if err := u.Insert(); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	response := &responseToken{
		Token: utils.GenerateToken(u.Id, u.Role),
	}
	return utils.SuccessRespond(c, http.StatusOK, response)
}

func LogIn(c echo.Context) error {
	email := c.QueryParam("email")
	pwd := c.QueryParam("pwd")
	u, err := model.CheckUser(email, pwd)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	response := &responseToken{
		Token: utils.GenerateToken(u.Id, u.Role),
	}
	return utils.SuccessRespond(c, http.StatusOK, response)
}

type responseUserInfo struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := model.GetUser(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	res := responseUserInfo{
		Id:       u.Id,
		Email:    u.Email,
		Username: u.Username,
	}
	return utils.SuccessRespond(c, http.StatusOK, res)
}

type receiveChangeInfo struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,max=20,min=6"`
	Pwd      string `json:"pwd" validate:"required,max=20,min=6"`
}

func ChangeInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	info := new(receiveChangeInfo)
	if err := c.Bind(info); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(info); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	pwdHash, err := utils.PwdHash(info.Pwd)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	err = model.Update(
		id,
		info.Email,
		info.Username,
		string(pwdHash),
	)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessRespond(c, http.StatusOK, nil)
}

type responseAllUser struct {
	Id         int       `json:"id"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	CreateTime time.Time `json:"createTime"`
	Role       bool      `json:"role"`
}

func GetAllUser(c echo.Context) error {
	users, err := model.SelectAllUser()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	usersInfo := make([]responseAllUser, len(users))
	for i, _ := range users {
		usersInfo[i].Id = users[i].Id
		usersInfo[i].Email = users[i].Email
		usersInfo[i].Username = users[i].Username
		usersInfo[i].CreateTime = users[i].CreateTime
		usersInfo[i].Role = users[i].Role
	}
	return utils.SuccessRespond(c, http.StatusOK, usersInfo)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := model.Check(id); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	err := model.DeleteUser(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessRespond(c, http.StatusOK, nil)
}
