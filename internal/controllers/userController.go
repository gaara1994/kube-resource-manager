package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kube-resource-manager/internal/dao"
	"kube-resource-manager/internal/dto"
	"kube-resource-manager/internal/errcodes"
	"kube-resource-manager/internal/response"
	"kube-resource-manager/utils/auth"
)

type UserController struct {
}

func (u *UserController) Get(c *gin.Context) {

}

func (u *UserController) Post(c *gin.Context) {
	req := dto.UserCreateRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.UserErrPost, errcodes.UserErrMsg[errcodes.UserErrPost])
		return
	}
	//查找用户名是否已存在
	user, err := dao.UserDao.GetByUsername(req.Username)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.UserErrPost, errcodes.UserErrMsg[errcodes.UserErrPost])
		return
	}
	fmt.Println(user.ID)
	if user.ID != 0 {
		//用户已经存在
		response.ErrorResponseWithMessage(c, errcodes.UserErrUserExisting, errcodes.UserErrMsg[errcodes.UserErrUserExisting])
		return
	}
	//密码加密
	hashPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.UserErrUserExisting, errcodes.UserErrMsg[errcodes.UserErrUserExisting])
		return
	}
	user.Username = req.Username
	user.PasswordHash = hashPassword

	err = dao.UserDao.Save(user)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.UserErrUserExisting, errcodes.UserErrMsg[errcodes.UserErrUserExisting])
		return
	}

	response.SuccessResponseWithMessage(c, nil)
}

func (u *UserController) Put(c *gin.Context) {

}
func (u *UserController) Delete(c *gin.Context) {

}
