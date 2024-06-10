package controllers

import (
	"github.com/gin-gonic/gin"
	"kube-resource-manager/internal/response"
	"kube-resource-manager/pkg/logger"
	"net/http"
	"strconv"
)

type KubernetesClusterController struct {
}

func (k *KubernetesClusterController) GET(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Log.Error("id参数错误")
		c.JSON(http.StatusBadRequest, response.ErrorResponse(response.ClusterErrGet, response.ClusterErrMsg[response.ClusterErrGet]))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(id))
}
