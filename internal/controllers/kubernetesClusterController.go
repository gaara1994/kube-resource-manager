package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kube-resource-manager/internal/dao"
	"kube-resource-manager/internal/db/models"
	dto "kube-resource-manager/internal/dto"
	"kube-resource-manager/internal/errcodes"
	"kube-resource-manager/internal/response"
	"strconv"
	"strings"
)

type KubernetesClusterController struct {
}

func (k *KubernetesClusterController) GET(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ClusterErrGet, errcodes.ClusterErrMsg[errcodes.ClusterErrGet])
		return
	}
	//查库
	data, err := dao.KubernetesClusterDao.GET(uint(id))
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ClusterErrGet, errcodes.ClusterErrMsg[errcodes.ClusterErrGet])
		return
	}
	response.SuccessResponseWithMessage(c, data)
}

func (k *KubernetesClusterController) POST(c *gin.Context) {
	req := dto.PostClusterRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ClusterErrPost, errcodes.ClusterErrMsg[errcodes.ClusterErrPost])
		return
	}

	//校验

	//赋值
	clusterModel := models.KubernetesCluster{
		ClusterName: req.ClusterName,
		APIEndpoint: req.APIEndpoint,
		KubeConfig:  req.KubeConfig,
		Version:     req.Version,
		Status:      models.KubernetesClusterStatus(req.Status),
		Description: req.Description,
	}
	//入库
	err = dao.KubernetesClusterDao.Save(&clusterModel)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ClusterErrPost, errcodes.ClusterErrMsg[errcodes.ClusterErrPost])
		return
	}
	response.SuccessResponseWithMessage(c, 1)
}

func (k *KubernetesClusterController) PUT(c *gin.Context) {
	req := dto.PostClusterRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ClusterErrPost, errcodes.ClusterErrMsg[errcodes.ClusterErrPost])
		return
	}

	//校验

	//赋值
	clusterModel := models.KubernetesCluster{
		Model: gorm.Model{
			ID: req.ID,
		},
		ClusterName: req.ClusterName,
		APIEndpoint: req.APIEndpoint,
		KubeConfig:  req.KubeConfig,
		Version:     req.Version,
		Status:      models.KubernetesClusterStatus(req.Status),
		Description: req.Description,
	}
	//入库
	err = dao.KubernetesClusterDao.Save(&clusterModel)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ClusterErrPost, errcodes.ClusterErrMsg[errcodes.ClusterErrPost])
		return
	}
	response.SuccessResponseWithMessage(c, 1)
}

func (k *KubernetesClusterController) DELETE(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ClusterErrDelete, errcodes.ClusterErrMsg[errcodes.ClusterErrDelete])
		return
	}
	//删除
	err = dao.KubernetesClusterDao.DELETE(uint(id))
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ClusterErrDelete, errcodes.ClusterErrMsg[errcodes.ClusterErrDelete])
		return
	}
	response.SuccessResponseWithMessage(c, nil)
}

// QueryList 根据多种查询条件查询列表
func (k *KubernetesClusterController) QueryList(c *gin.Context) {
	clusterName := c.Query("cluster_name")
	description := c.Query("description")
	status := c.Query("status")

	queryConditions := make([]string, 0)
	queryValues := make([]interface{}, 0)

	// cluster_name 的处理
	if clusterName != "" {
		keyword := "%" + clusterName + "%"
		queryConditions = append(queryConditions, "cluster_name LIKE ?")
		queryValues = append(queryValues, keyword)
	}

	// description 的处理
	if description != "" {
		keyword := "%" + description + "%"
		if len(queryConditions) > 0 { // 非第一个条件时添加AND
			queryConditions = append(queryConditions, "AND description LIKE ?")
		} else {
			queryConditions = append(queryConditions, "description LIKE ?")
		}
		queryValues = append(queryValues, keyword)
	}

	// status 的处理
	if status != "" {
		if len(queryConditions) > 0 { // 非第一个条件时添加AND
			queryConditions = append(queryConditions, "AND status = ?")
		} else {
			queryConditions = append(queryConditions, "status = ?")
		}
		queryValues = append(queryValues, status)
	}

	// 构造查询字符串
	queryStr := strings.Join(queryConditions, " ")
	date, err := dao.KubernetesClusterDao.List(queryStr, queryValues)
	fmt.Println(date)
	fmt.Println(err)
	return
	// 根据存在的条件构造查询
	if len(queryConditions) > 0 {
		queryStr := strings.Join(queryConditions, " AND ")
		date, err := dao.KubernetesClusterDao.List(queryStr, queryValues)
		if err != nil {

		}
		fmt.Println(date)
	}
}
