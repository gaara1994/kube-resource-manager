package routes

import "kube-resource-manager/internal/controllers"

// 创建控制器实例
var (
	KubernetesClusterCtl = controllers.KubernetesClusterController{}
	UserCtl              = controllers.UserController{}
	LoginCtl             = controllers.LoginController{}
)
