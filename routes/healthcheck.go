package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

func HealthCheck(c *gin.Context) {
	// 这里可以添加具体的健康检查逻辑，比如检查数据库连接、依赖服务等
	// 下面是一个简单的示例，始终返回200 OK
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}