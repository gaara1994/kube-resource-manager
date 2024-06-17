package models

import (
	"gorm.io/gorm"
)

// KubernetesResourceType 代表Kubernetes资源类型的GORM模型
type KubernetesResourceType struct {
	gorm.Model
	ResourceName string `gorm:"column:resource_name;not null;comment:资源类型名称"` // 资源类型名称，非空
	Description  string `gorm:"column:description;type:text;comment:描述信息"`    // 描述信息，使用TEXT类型存储大文本
}
