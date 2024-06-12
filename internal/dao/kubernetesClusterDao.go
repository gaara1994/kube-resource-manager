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

func (k *kubernetesCluster) List(clusterName string, description string, status string, page int, pageSize int) ([]models.KubernetesCluster, error) {
	var clusters []models.KubernetesCluster
	offset := (page - 1) * pageSize
	query := db.DB.Model(&models.KubernetesCluster{})

	if clusterName != "" {
		query = query.Where("cluster_name like ?", "%"+clusterName+"%")
	}
	if description != "" {
		query = query.Where("description like ?", "%"+description+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Offset(offset).Limit(pageSize).Find(&clusters).Error
	return clusters, err
}
