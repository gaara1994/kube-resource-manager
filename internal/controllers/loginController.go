package controllers

import (
	"github.com/gin-gonic/gin"
	"kube-resource-manager/internal/dao"
	"kube-resource-manager/internal/dto"
	"kube-resource-manager/internal/errcodes"
	"kube-resource-manager/internal/response"
	"kube-resource-manager/utils/auth"
)

type LoginController struct {
}

func (l *LoginController) Login(c *gin.Context) {
	req := dto.UserLoginRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.UserErrLogin, errcodes.UserErrMsg[errcodes.UserErrLogin])
		return
	}
	//查询用户
	user, err := dao.UserDao.GetByUsername(req.Username)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.UserErrLogin, errcodes.UserErrMsg[errcodes.UserErrLogin])
		return
	}
	if user.ID == 0 {
		response.HandleErrorAndRespond(c, err, errcodes.UserErrUserNotExisting, errcodes.UserErrMsg[errcodes.UserErrUserNotExisting])
		return
	}
	//对比用户密码
	if !auth.ComparePasswords(user.PasswordHash, req.Password) {
		response.HandleErrorAndRespond(c, err, errcodes.UserErrLoginPassword, errcodes.UserErrMsg[errcodes.UserErrLoginPassword])
		return
	}
	token, err := generateToken(user)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.UserErrLoginPassword, errcodes.UserErrMsg[errcodes.UserErrLoginPassword])
		return
	}
	response.SuccessResponseWithMessage(c, token)
}
