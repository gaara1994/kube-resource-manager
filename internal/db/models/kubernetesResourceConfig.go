package models

import (
	"gorm.io/gorm"
)

// KubernetesResourceConfig 代表Kubernetes资源配置的GORM模型
type KubernetesResourceConfig struct {
	BaseModel
	ResourceTypeID uint   `gorm:"column:resource_type_id;not null;comment:资源类型ID"`         // 资源类型ID，与k8s_resource_types表的外键关联
	YamlContent    string `gorm:"column:yaml_content;type:text;not null;comment:YAML配置内容"` // YAML配置内容，使用TEXT类型存储大文本，非空
	Description    string `gorm:"column:description;type:text;comment:配置描述信息"`             // 配置描述信息，可选
}

// BeforeCreate 钩子，在创建记录前执行
func (*KubernetesResourceConfig) BeforeCreate(tx *gorm.DB) error {
	// 这里可以添加在保存之前需要执行的逻辑，例如验证YAML格式等
	return nil
}

func (*KubernetesResourceConfig) TableName() string {
	return "kubernetes_resource_config"
}
