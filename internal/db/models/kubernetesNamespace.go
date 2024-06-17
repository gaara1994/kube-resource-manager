package models

// KubernetesNamespace 代表Kubernetes集群的GORM模型
type KubernetesNamespace struct {
	BaseModel
	Name        string   `gorm:"column:name;not null;comment:命名空间名称"` // 命名空间名称，唯一且非空
	ClusterID   uint     `gorm:"column:cluster_id;not null;comment:集群id"`
	Description string   `gorm:"column:description;comment:描述"`
	Tags        []string `gorm:"serializer:json"` // 标签列表，用JSON序列化存储
}

func (kn *KubernetesNamespace) TableName() string {
	return "kubernetes_namespaces"
}
