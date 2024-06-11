package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kube-resource-manager/config"
	"kube-resource-manager/internal/db/models"
)

var DB *gorm.DB

func InitDb() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		config.Config.Postgres.Host, config.Config.Postgres.User, config.Config.Postgres.Password, config.Config.Postgres.DBName, config.Config.Postgres.Port, config.Config.Postgres.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	// 设置最大打开的连接数
	sqlDB.SetMaxOpenConns(config.Config.Postgres.MaxOpen)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.Config.Postgres.MaxIdle)

	//自动迁移
	DB = db
	err = db.AutoMigrate(&models.KubernetesCluster{}, &models.K8sNamespace{}, &models.K8sResourceType{}, &models.K8sResourceConfig{}, &models.Permission{}, &models.Role{}, &models.RolePermission{}, &models.User{}, &models.UserRole{})
	if err != nil {
		return err
	}
	return nil
}
