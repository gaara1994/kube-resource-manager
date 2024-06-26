package main

import (
	"github.com/gin-gonic/gin"
	"kube-resource-manager/cmd"
	"kube-resource-manager/config"
	"kube-resource-manager/internal/db"
	"kube-resource-manager/pkg/logger"
	"kube-resource-manager/routes"
	"log"
	"strconv"
)

func main() {
	cmd.Execute()

	//初始化配置
	config.InitConfig()

	//初始化日志
	err := logger.InitLogger()
	if err != nil {
		log.Fatalf(err.Error())
	}

	//初始化数据库
	err = db.InitDb()
	if err != nil {
		logger.Log.Error(err.Error())
	}

	//设置模式
	gin.SetMode(config.Config.Server.Mode)
	// 启动HTTP服务器
	r := gin.Default()
	routes.InitRoutes(r)
	if err := r.Run(":8080"); err != nil {
		logger.Log.Error(err.Error())
	}
	if err := r.Run(":" + strconv.Itoa(config.Config.Server.Port)); err != nil {
		logger.Log.Error(err.Error())
	}

}
