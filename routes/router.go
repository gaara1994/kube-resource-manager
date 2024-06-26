package routes

import (
	"kube-resource-manager/utils/auth"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRoutes initializes all the necessary routes for the Gin engine.
func InitRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.Get("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	// 添加 Prometheus metrics 路由
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// 注册健康检查路由，通常使用 /healthz 或 /ping
	r.GET("/healthz", HealthCheck)

	api := r.Group("api")
	api.Use(auth.AuthMiddleware())
	{
		v1 := api.Group("v1")
		{
			//服务器集群 Server cluster
			v1.GET("/cluster/:id", KubernetesClusterCtl.GET)
			v1.POST("/cluster", KubernetesClusterCtl.POST)
			v1.PUT("/cluster", KubernetesClusterCtl.PUT)
			v1.DELETE("/cluster/:id", KubernetesClusterCtl.DELETE)
			v1.GET("/cluster/list", KubernetesClusterCtl.QueryList)

			//名称空间
			v1.GET("/namespace/:id", KubernetesNamespaceCtl.GET)
			v1.POST("/namespace", KubernetesNamespaceCtl.POST)
			v1.PUT("/namespace", KubernetesNamespaceCtl.PUT)
			v1.DELETE("/namespace/:id", KubernetesNamespaceCtl.DELETE)
			v1.GET("/namespace/list", KubernetesNamespaceCtl.QueryList)

			//资源分类
			v1.GET("/resource/type/:id", KubernetesResourceTypeCtl.GET)
			v1.POST("/resource/type", KubernetesResourceTypeCtl.POST)
			v1.PUT("/resource/type", KubernetesResourceTypeCtl.PUT)
			v1.DELETE("/resource/type/:id", KubernetesResourceTypeCtl.DELETE)
			v1.GET("/resource/type/list", KubernetesResourceTypeCtl.QueryList)

			//资源文件
			v1.GET("/resource/config/:id", KubernetesResourceConfigCtl.GET)
			v1.POST("/resource/config", KubernetesResourceConfigCtl.POST)
			v1.PUT("/resource/config", KubernetesResourceConfigCtl.PUT)
			v1.DELETE("/resource/config/:id", KubernetesResourceConfigCtl.DELETE)
			v1.GET("/resource/config/list", KubernetesResourceConfigCtl.QueryList)

			//用户
			v1.POST("/user", UserCtl.Post)

		}
	}

	//登录
	r.POST("/login", LoginCtl.Login)

}
