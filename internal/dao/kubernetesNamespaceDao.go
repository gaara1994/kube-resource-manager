package dao

import (
	"kube-resource-manager/internal/db"
	"kube-resource-manager/internal/db/models"
)

var KubernetesNamespaceDao *KubernetesNamespace

type KubernetesNamespace struct {
}

func (kn *KubernetesNamespace) Get(id uint) (*models.KubernetesNamespace, error) {
	m := new(models.KubernetesNamespace)
	return m, db.DB.Table(m.TableName()).Where("id = ?", id).Find(m).Error
}

func (kn *KubernetesNamespace) Save(m *models.KubernetesNamespace) error {
	return db.DB.Table(m.TableName()).Save(m).Error
}

func (kn *KubernetesNamespace) Delete(id uint) error {
	m := new(models.KubernetesNamespace)
	return db.DB.Table(m.TableName()).Where("id = ?", id).Delete(m).Error
}

func (kn *KubernetesNamespace) List(clusterName string, description string, status string, page int, pageSize int) ([]models.KubernetesNamespace, error) {
	var clusters []models.KubernetesNamespace
	offset := (page - 1) * pageSize
	query := db.DB.Model(&models.KubernetesNamespace{})

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
