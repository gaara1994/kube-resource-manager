package models

import (
	"gorm.io/gorm"
)

// K8sNamespace 代表Kubernetes集群的GORM模型
type K8sNamespace struct {
	gorm.Model
	Name        string   `gorm:"column:name;unique;not null;comment:命名空间名称"` // 命名空间名称，唯一且非空
	ClusterID   uint     `gorm:"column:cluster_id;not null;comment:集群id"`
	Description string   `gorm:"column:description;comment:描述"`
	Tags        []string `gorm:"serializer:json"` // 标签列表，用JSON序列化存储
}
