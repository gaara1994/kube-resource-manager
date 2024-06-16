package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"kube-resource-manager/internal/dao"
	"kube-resource-manager/internal/db/models"
	"kube-resource-manager/internal/dto"
	"kube-resource-manager/internal/errcodes"
	"kube-resource-manager/internal/response"
	"kube-resource-manager/utils/auth"
	"time"
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

var jwtSecret = []byte("kube-resource-manager")

func generateToken(user *models.User) (string, error) {
	claims := createClaims(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func createClaims(user *models.User) jwt.MapClaims {
	return jwt.MapClaims{
		"UserId":   user.ID,
		"Username": user.Username,
		"Exp":      time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为24小时后
	}
}
