package dto

import requests "kube-resource-manager/internal/request"

// GetClusterListRequest 用于获取集群列表的请求参数
type GetClusterListRequest struct {
	requests.Pagination        // 嵌入分页结构体
	SearchTerm          string `json:"search_term" form:"search_term"` // 可选的搜索关键词
}

// PostClusterRequest 用于获取集群的请求参数
type PostClusterRequest struct {
	ID          uint   `json:"id"`
	ClusterName string `json:"cluster_name" binding:"required,max=256"`         // 确保请求体中的字段符合预期
	APIEndpoint string `json:"api_endpoint" binding:"required,url"`             // 对endpoint进行URL格式验证
	KubeConfig  string `json:"kube_config,omitempty"`                           // 可选字段
	Version     string `json:"version" binding:"required,max=50"`               // 版本验证
	Status      string `json:"status" binding:"omitempty,oneof=Running Paused"` // 状态验证
	Description string `json:"description" binding:"omitempty,max=50"`          // 描述验证
}
