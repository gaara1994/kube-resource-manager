package dao

import (
	"kube-resource-manager/internal/db"
	"kube-resource-manager/internal/db/models"
)

var KubernetesResourceConfigDao *KubernetesResourceConfig

type KubernetesResourceConfig struct {
}

func (krt *KubernetesResourceConfig) Get(id uint) (*models.KubernetesResourceConfig, error) {
	m := new(models.KubernetesResourceConfig)
	return m, db.DB.Table(m.TableName()).Where("id = ?", id).Find(m).Error
}

func (krt *KubernetesResourceConfig) Save(m *models.KubernetesResourceConfig) error {
	return db.DB.Table(m.TableName()).Save(m).Error
}

func (krt *KubernetesResourceConfig) Delete(id uint) error {
	m := new(models.KubernetesResourceConfig)
	return db.DB.Table(m.TableName()).Where("id = ?", id).Delete(m).Error
}

func (krt *KubernetesResourceConfig) List(clusterName string, description string, status string, page int, pageSize int) ([]models.KubernetesResourceConfig, error) {
	var clusters []models.KubernetesResourceConfig
	offset := (page - 1) * pageSize
	query := db.DB.Model(&models.KubernetesResourceConfig{})

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
