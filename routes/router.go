package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// InitRoutes initializes all the necessary routes for the Gin engine.
func InitRoutes(r *gin.Engine) {
	// 添加 Prometheus metrics 路由
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// 注册健康检查路由，通常使用 /healthz 或 /ping
	r.GET("/healthz", HealthCheck)


}

