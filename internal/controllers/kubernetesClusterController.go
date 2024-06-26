package controllers

import (
	"github.com/gin-gonic/gin"
	"kube-resource-manager/internal/dao"
	"kube-resource-manager/internal/db/models"
	dto "kube-resource-manager/internal/dto"
	"kube-resource-manager/internal/errcodes"
	"kube-resource-manager/internal/response"
	"strconv"
)

type KubernetesClusterController struct {
}

// GET PathParamsExample godoc
//
//	@Summary		path params example
//	@Description	path params
//	@Tags			example
//	@Accept			json
//	@Produce		plain
//	@Param			id			path		int		true	"id"
//	@Success		200			{string}	string	"answer"
//	@Failure		400			{string}	string	"ok"
//	@Failure		404			{string}	string	"ok"
//	@Failure		500			{string}	string	"ok"
//	@Router			/api/v1/cluster/{id}/ [get]
func (k *KubernetesClusterController) GET(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ClusterErrGet, errcodes.ClusterErrMsg[errcodes.ClusterErrGet])
		return
	}
	//查库
	data, err := dao.KubernetesClusterDao.Get(uint(id))
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ClusterErrGet, errcodes.ClusterErrMsg[errcodes.ClusterErrGet])
		return
	}
	if data.ID == 0 {
		response.SuccessResponseWithMessage(c, nil)
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
	response.SuccessResponseWithMessage(c, nil)
}

func (k *KubernetesClusterController) PUT(c *gin.Context) {
	req := models.KubernetesCluster{}
	err := c.BindJSON(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ClusterErrPost, errcodes.ClusterErrMsg[errcodes.ClusterErrPost])
		return
	}

	//校验

	//赋值
	//入库
	err = dao.KubernetesClusterDao.Save(&req)
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
	err = dao.KubernetesClusterDao.Delete(uint(id))
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
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	list, err := dao.KubernetesClusterDao.List(clusterName, description, status, page, pageSize)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ClusterErrQueryList, errcodes.ClusterErrMsg[errcodes.ClusterErrQueryList])
		return
	}
	response.SuccessResponseWithMessage(c, list)
}
