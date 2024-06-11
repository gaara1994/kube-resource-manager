package dao

import (
	"kube-resource-manager/internal/db"
	"kube-resource-manager/internal/db/models"
)

var KubernetesClusterDao *kubernetesCluster

type kubernetesCluster struct {
}

func (k *kubernetesCluster) GET(id uint) (*models.KubernetesCluster, error) {
	m := new(models.KubernetesCluster)
	return m, db.DB.Table(m.TableName()).Where("id = ?", id).Find(m).Error
}

func (k *kubernetesCluster) Save(m *models.KubernetesCluster) error {
	return db.DB.Table(m.TableName()).Save(m).Error
}

func (k *kubernetesCluster) DELETE(id uint) error {
	m := new(models.KubernetesCluster)
	return db.DB.Table(m.TableName()).Where("id = ?", id).Delete(m).Error
}

func (k *kubernetesCluster) List(queryStr string, queryValues []interface{}) ([]models.KubernetesCluster, error) {
	m := new(models.KubernetesCluster)
	var data []models.KubernetesCluster
	err := db.DB.Table(m.TableName()).Where(queryStr, queryValues).Find(&data).Error
	return data, err
}
