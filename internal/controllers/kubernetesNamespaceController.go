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

type KubernetesNamespaceController struct {
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
func (k *KubernetesNamespaceController) GET(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.NamespaceErrGet, errcodes.NamespaceErrMsg[errcodes.NamespaceErrGet])
		return
	}
	//查库
	data, err := dao.KubernetesNamespaceDao.Get(uint(id))
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.NamespaceErrGet, errcodes.NamespaceErrMsg[errcodes.NamespaceErrGet])
		return
	}
	if data.ID == 0 {
		response.HandleErrorAndRespond(c, nil, errcodes.NamespaceErrGet, errcodes.NamespaceErrMsg[errcodes.NamespaceErrGet])
		return
	}
	response.SuccessResponseWithMessage(c, data)
}

func (k *KubernetesNamespaceController) POST(c *gin.Context) {
	req := models.KubernetesNamespace{}
	err := c.BindJSON(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.NamespaceErrPost, errcodes.NamespaceErrMsg[errcodes.NamespaceErrPost])
		return
	}

	//校验

	//赋值

	//入库
	fmt.Println("req=", req)
	err = dao.KubernetesNamespaceDao.Save(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.NamespaceErrPost, errcodes.NamespaceErrMsg[errcodes.NamespaceErrPost])
		return
	}
	response.SuccessResponseWithMessage(c, nil)
}

func (k *KubernetesNamespaceController) PUT(c *gin.Context) {
	req := models.KubernetesNamespace{}
	err := c.BindJSON(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.NamespaceErrPost, errcodes.NamespaceErrMsg[errcodes.NamespaceErrPost])
		return
	}

	//校验

	//赋值

	//入库
	err = dao.KubernetesNamespaceDao.Save(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.NamespaceErrPost, errcodes.NamespaceErrMsg[errcodes.NamespaceErrPost])
		return
	}
	response.SuccessResponseWithMessage(c, 1)
}

func (k *KubernetesNamespaceController) DELETE(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.NamespaceErrDelete, errcodes.NamespaceErrMsg[errcodes.NamespaceErrDelete])
		return
	}
	//删除
	err = dao.KubernetesNamespaceDao.Delete(uint(id))
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.NamespaceErrDelete, errcodes.NamespaceErrMsg[errcodes.NamespaceErrDelete])
		return
	}
	response.SuccessResponseWithMessage(c, nil)
}

// QueryList 根据多种查询条件查询列表
func (k *KubernetesNamespaceController) QueryList(c *gin.Context) {
	name := c.Query("name")
	description := c.Query("description")
	status := c.Query("status")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	list, err := dao.KubernetesNamespaceDao.List(name, description, status, page, pageSize)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.NamespaceErrQueryList, errcodes.NamespaceErrMsg[errcodes.NamespaceErrQueryList])
		return
	}
	response.SuccessResponseWithMessage(c, list)
}
