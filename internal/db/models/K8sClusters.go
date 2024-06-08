package models

import (
	"gorm.io/gorm"
)

// KubernetesClusterStatus 定义了集群状态的枚举类型
type KubernetesClusterStatus string

const (
	Running KubernetesClusterStatus = "Running"
	Stopped KubernetesClusterStatus = "Stopped"
	Error   KubernetesClusterStatus = "Error"
)

// KubernetesCluster 代表Kubernetes集群的GORM模型
type KubernetesCluster struct {
	gorm.Model
	ClusterName string                  `gorm:"unique;not null"`     // 设置ClusterName为唯一且非空
	APIEndpoint string                  `gorm:"not null"`            // API Endpoint必须填写
	KubeConfig  string                  `gorm:"type:text"`           // 可选，Kubeconfig的Base64编码
	ClusterCIDR string                  `gorm:"column:cluster_cidr"` // 集群CIDR
	ServiceCIDR string                  `gorm:"column:service_cidr"` // 服务CIDR
	Version     string                  `gorm:"size:50"`             // Kubernetes版本
	Status      KubernetesClusterStatus `gorm:"default:'Running'"`   // 集群状态，默认为Running
	Description string                  `gorm:"type:text"`           // 集群描述
}
