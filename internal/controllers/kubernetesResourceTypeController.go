package controllers

import (
	"fmt"
	"kube-resource-manager/internal/dao"
	"kube-resource-manager/internal/db/models"
	"kube-resource-manager/internal/errcodes"
	"kube-resource-manager/internal/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KubernetesResourceTypeController struct {
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
func (k *KubernetesResourceTypeController) GET(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceTypeErrGet, errcodes.ResourceTypeErrMsg[errcodes.ResourceTypeErrGet])
		return
	}
	//查库
	data, err := dao.KubernetesResourceTypeDao.Get(uint(id))
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceTypeErrGet, errcodes.ResourceTypeErrMsg[errcodes.ResourceTypeErrGet])
		return
	}
	if data.ID == 0 {
		response.HandleErrorAndRespond(c, nil, errcodes.ResourceTypeErrGet, errcodes.ResourceTypeErrMsg[errcodes.ResourceTypeErrGet])
		return
	}
	response.SuccessResponseWithMessage(c, data)
}

func (k *KubernetesResourceTypeController) POST(c *gin.Context) {
	req := models.KubernetesResourceType{}
	err := c.BindJSON(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceTypeErrPost, errcodes.ResourceTypeErrMsg[errcodes.ResourceTypeErrPost])
		return
	}

	//校验

	//赋值

	//入库
	fmt.Println("req=", req.ResourceName)
	fmt.Println("req=", req.Description)
	err = dao.KubernetesResourceTypeDao.Save(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceTypeErrPost, errcodes.ResourceTypeErrMsg[errcodes.ResourceTypeErrPost])
		return
	}
	response.SuccessResponseWithMessage(c, nil)
}

func (k *KubernetesResourceTypeController) PUT(c *gin.Context) {
	req := models.KubernetesResourceType{}
	err := c.BindJSON(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceTypeErrPost, errcodes.ResourceTypeErrMsg[errcodes.ResourceTypeErrPost])
		return
	}

	//校验

	//赋值

	//入库
	err = dao.KubernetesResourceTypeDao.Save(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceTypeErrPost, errcodes.ResourceTypeErrMsg[errcodes.ResourceTypeErrPost])
		return
	}
	response.SuccessResponseWithMessage(c, 1)
}

func (k *KubernetesResourceTypeController) DELETE(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceTypeErrDelete, errcodes.ResourceTypeErrMsg[errcodes.ResourceTypeErrDelete])
		return
	}
	//删除
	err = dao.KubernetesResourceTypeDao.Delete(uint(id))
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceTypeErrDelete, errcodes.ResourceTypeErrMsg[errcodes.ResourceTypeErrDelete])
		return
	}
	response.SuccessResponseWithMessage(c, nil)
}

// QueryList 根据多种查询条件查询列表
func (k *KubernetesResourceTypeController) QueryList(c *gin.Context) {
	name := c.Query("name")
	description := c.Query("description")
	status := c.Query("status")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	list, err := dao.KubernetesResourceTypeDao.List(name, description, status, page, pageSize)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceTypeErrQueryList, errcodes.ResourceTypeErrMsg[errcodes.ResourceTypeErrQueryList])
		return
	}
	response.SuccessResponseWithMessage(c, list)
}
