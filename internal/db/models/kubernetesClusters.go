package models

// KubernetesClusterStatus 定义了集群状态的枚举类型
type KubernetesClusterStatus string

const (
	Running KubernetesClusterStatus = "Running"
	Stopped KubernetesClusterStatus = "Stopped"
	Error   KubernetesClusterStatus = "Error"
)

// KubernetesCluster 代表Kubernetes集群的GORM模型
type KubernetesCluster struct {
	BaseModel
	ClusterName string                  `json:"cluster_name" gorm:"column:cluster_name;size:256;unique;not null;comment:集群名称"` // 设置ClusterName为唯一且非空
	APIEndpoint string                  `json:"api_endpoint" gorm:"column:api_endpoint;size:256;comment:APIEndpoint"`          // API Endpoint必须填写
	KubeConfig  string                  `json:"kube_config" gorm:"column:kube_config;comment:KubeConfig"`                     // 可选，KubeConfig的Base64编码
	Version     string                  `json:"version" gorm:"column:version;size:50;comment:k8s版本"`                      // Kubernetes版本
	Status      KubernetesClusterStatus `json:"status" gorm:"column:status;size:50;default:'Running';comment:集群状态"`      // 集群状态，默认为Running
	Description string                  `json:"description" gorm:"column:description;type:string;size:50;comment:集群描述"`       // 集群描述
}

func (*KubernetesCluster) TableName() string {
	return "kubernetes_clusters"
}
