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

	api := r.Group("api")
	{
		v1 := api.Group("v1")
		{
			//服务器集群 Server cluster
			v1.GET("/cluster/:id", KubernetesClusterCtl.GET)
			v1.POST("/cluster", KubernetesClusterCtl.POST)
			v1.PUT("/cluster", KubernetesClusterCtl.PUT)
			v1.DELETE("/cluster/:id", KubernetesClusterCtl.DELETE)
			v1.GET("/cluster/list", KubernetesClusterCtl.QueryList)

			//资源分类
			v1.GET("/resource/type")
			v1.POST("/resource/type")
			v1.PUT("/resource/type")
			v1.DELETE("/resource/type")
			v1.GET("/resource/type/list")

			//资源文件
			v1.GET("/resource/config")
			v1.POST("/resource/config")
			v1.PUT("/resource/config")
			v1.DELETE("/resource/config")
			v1.GET("/resource/config/list")
		}
	}

}
