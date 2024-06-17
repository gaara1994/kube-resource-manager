package dao

import (
	"kube-resource-manager/internal/db"
	"kube-resource-manager/internal/db/models"
)

var KubernetesResourceTypeDao *KubernetesResourceType

type KubernetesResourceType struct {
}

func (krt *KubernetesResourceType) Get(id uint) (*models.KubernetesResourceType, error) {
	m := new(models.KubernetesResourceType)
	return m, db.DB.Table(m.TableName()).Where("id = ?", id).Find(m).Error
}

func (krt *KubernetesResourceType) Save(m *models.KubernetesResourceType) error {
	return db.DB.Table(m.TableName()).Save(m).Error
}

func (krt *KubernetesResourceType) Delete(id uint) error {
	m := new(models.KubernetesResourceType)
	return db.DB.Table(m.TableName()).Where("id = ?", id).Delete(m).Error
}

func (krt *KubernetesResourceType) List(clusterName string, description string, status string, page int, pageSize int) ([]models.KubernetesResourceType, error) {
	var clusters []models.KubernetesResourceType
	offset := (page - 1) * pageSize
	query := db.DB.Model(&models.KubernetesResourceType{})

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
