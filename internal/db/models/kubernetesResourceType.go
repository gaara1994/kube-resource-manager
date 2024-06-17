package models

// KubernetesResourceType 代表Kubernetes资源类型的GORM模型
type KubernetesResourceType struct {
	BaseModel
	ResourceName string `json:"resource_name" gorm:"column:resource_name;not null;comment:资源类型名称"` // 资源类型名称，非空
	Description  string `json:"description" gorm:"column:description;type:text;comment:描述信息"`      // 描述信息，使用TEXT类型存储大文本
}

func (*KubernetesResourceType) TableName() string {
	return "kubernetes_resource_type"
}
