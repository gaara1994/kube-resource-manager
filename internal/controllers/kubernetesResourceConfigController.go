package controllers

import (
	"kube-resource-manager/internal/dao"
	"kube-resource-manager/internal/db/models"
	"kube-resource-manager/internal/errcodes"
	"kube-resource-manager/internal/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KubernetesResourceConfigController struct {
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
func (k *KubernetesResourceConfigController) GET(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceConfigErrGet, errcodes.ResourceConfigErrMsg[errcodes.ResourceConfigErrGet])
		return
	}
	//查库
	data, err := dao.KubernetesResourceConfigDao.Get(uint(id))
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceConfigErrGet, errcodes.ResourceConfigErrMsg[errcodes.ResourceConfigErrGet])
		return
	}
	if data.ID == 0 {
		response.HandleErrorAndRespond(c, nil, errcodes.ResourceConfigErrGet, errcodes.ResourceConfigErrMsg[errcodes.ResourceConfigErrGet])
		return
	}
	response.SuccessResponseWithMessage(c, data)
}

func (k *KubernetesResourceConfigController) POST(c *gin.Context) {
	req := models.KubernetesResourceConfig{}
	err := c.BindJSON(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceConfigErrPost, errcodes.ResourceConfigErrMsg[errcodes.ResourceConfigErrPost])
		return
	}

	//校验

	//赋值

	//入库
	err = dao.KubernetesResourceConfigDao.Save(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceConfigErrPost, errcodes.ResourceConfigErrMsg[errcodes.ResourceConfigErrPost])
		return
	}
	response.SuccessResponseWithMessage(c, nil)
}

func (k *KubernetesResourceConfigController) PUT(c *gin.Context) {
	req := models.KubernetesResourceConfig{}
	err := c.BindJSON(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceConfigErrPost, errcodes.ResourceConfigErrMsg[errcodes.ResourceConfigErrPost])
		return
	}

	//校验

	//赋值

	//入库
	err = dao.KubernetesResourceConfigDao.Save(&req)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceConfigErrPost, errcodes.ResourceConfigErrMsg[errcodes.ResourceConfigErrPost])
		return
	}
	response.SuccessResponseWithMessage(c, 1)
}

func (k *KubernetesResourceConfigController) DELETE(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceConfigErrDelete, errcodes.ResourceConfigErrMsg[errcodes.ResourceConfigErrDelete])
		return
	}
	//删除
	err = dao.KubernetesResourceConfigDao.Delete(uint(id))
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceConfigErrDelete, errcodes.ResourceConfigErrMsg[errcodes.ResourceConfigErrDelete])
		return
	}
	response.SuccessResponseWithMessage(c, nil)
}

// QueryList 根据多种查询条件查询列表
func (k *KubernetesResourceConfigController) QueryList(c *gin.Context) {
	name := c.Query("name")
	description := c.Query("description")
	status := c.Query("status")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	list, err := dao.KubernetesResourceConfigDao.List(name, description, status, page, pageSize)
	if err != nil {
		response.HandleErrorAndRespond(c, err, errcodes.ResourceConfigErrQueryList, errcodes.ResourceConfigErrMsg[errcodes.ResourceConfigErrQueryList])
		return
	}
	response.SuccessResponseWithMessage(c, list)
}
