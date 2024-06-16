package dao

import (
	"kube-resource-manager/internal/db"
	"kube-resource-manager/internal/db/models"
)

var UserDao *user

type user struct {
}

func (u *user) Get(id uint) (*models.User, error) {
	m := new(models.User)
	return m, db.DB.Table(m.TableName()).Where("id = ?", id).Find(m).Error
}

func (u *user) GetByUsername(Username string) (*models.User, error) {
	m := new(models.User)
	return m, db.DB.Table(m.TableName()).Where("user_name = ?", Username).Find(m).Error
}

func (u *user) Save(m *models.User) error {
	return db.DB.Table(m.TableName()).Save(m).Error
}

func (u *user) DELETE(id uint) error {
	m := new(models.User)
	return db.DB.Table(m.TableName()).Where("id = ?", id).Delete(m).Error
}

func (u *user) List(clusterName string, description string, status string, page int, pageSize int) ([]models.User, error) {
	var clusters []models.User
	offset := (page - 1) * pageSize
	query := db.DB.Model(&models.User{})

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
