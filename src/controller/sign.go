package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	model "src/model"
	"src/utils"
	"strconv"
)

type receiveSignUp struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

type responseToken struct {
	Token string `json:"token"`
}

func SignUP(c echo.Context) error {
	receive := new(receiveSignUp)
	if err := c.Bind(receive); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	pwdHash := utils.PwdHash(receive.Pwd)
	u := &model.User{
		Email:    receive.Email,
		Username: receive.Username,
		PwdHash:  pwdHash,
	}
	if err := u.Insert(); err != nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	response := &responseToken{
		Token: utils.GenerateToken(u.Id, u.Role),
	}
	return utils.SuccessRespond(c, http.StatusOK, response)
}

func LogIn(c echo.Context) error {
	email := c.QueryParam("email")
	pwd := c.QueryParam("pwd")
	pwdHash := utils.PwdHash(pwd)
	u, err := model.CheckUser(email, pwdHash)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if u == nil {
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
	Email    string `json:"email"`
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

func ChangeInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	info := new(receiveChangeInfo)
	if err := c.Bind(info); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	err := model.Update(
		id,
		info.Email,
		info.Username,
		utils.PwdHash(info.Pwd),
	)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	return utils.SuccessRespond(c, http.StatusOK, nil)
}

func GetAllUser(c echo.Context) error {
	users, err := model.SelectAllUser()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessRespond(c, http.StatusOK, users)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := model.DeleteUser(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	return utils.SuccessRespond(c, http.StatusOK, nil)
}
