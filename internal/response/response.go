package response

import (
	"github.com/gin-gonic/gin"
	"kube-resource-manager/internal/errcodes"
	"kube-resource-manager/pkg/logger"
	"net/http"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // 数据字段，根据需要传入具体类型或nil
}

// SuccessResponse 创建一个成功的响应
func SuccessResponse(code int, data interface{}) BaseResponse {
	return BaseResponse{
		Code:    code,
		Message: "success",
		Data:    data,
	}
}

// ErrorResponse 创建一个错误响应
func ErrorResponse(code int, message string) BaseResponse {
	return BaseResponse{
		Code:    code,
		Message: message,
	}
}

func ErrorResponseWithMessage(c *gin.Context, errorCode int, errorMessage string) {
	c.JSON(http.StatusBadRequest, ErrorResponse(errorCode, errorMessage))
}

// SuccessResponseWithMessage 创建并发送一个带有自定义消息的成功响应
func SuccessResponseWithMessage(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse(errcodes.SuccessCode, data))
}

func HandleErrorAndRespond(c *gin.Context, err error, errorCode int, errorMessage string) {
	if err != nil {
		logger.Log.Error(err.Error())
		ErrorResponseWithMessage(c, errorCode, errorMessage)
		c.Abort() // 确保后续操作不会被执行
	}
}
