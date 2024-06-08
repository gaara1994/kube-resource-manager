package models

import (
	"gorm.io/gorm"
	"time"
)

// K8sResourceType 代表Kubernetes资源类型的GORM模型
type K8sResourceType struct {
	gorm.Model
	ResourceName string    `gorm:"not null"`       // 资源名称，非空
	ClusterID    uint      `gorm:"not null"`       // 所属集群，非空
	Description  string    `gorm:"type:text"`      // 描述信息，使用TEXT类型存储大文本
	CreatedAt    time.Time `gorm:"autoCreateTime"` // 创建时间，自动设置
	UpdatedAt    time.Time `gorm:"autoUpdateTime"` // 更新时间，自动更新
}
