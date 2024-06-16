package db

import (
	"fmt"
	"kube-resource-manager/config"
	"kube-resource-manager/internal/db/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDb() error {
	// 解析日志级别
	var logLevel logger.LogLevel
	switch config.Config.Postgres.Debug {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		log.Fatal("Invalid log level specified in the configuration.")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		config.Config.Postgres.Host, config.Config.Postgres.User, config.Config.Postgres.Password, config.Config.Postgres.DBName, config.Config.Postgres.Port, config.Config.Postgres.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   logger.Default.LogMode(logLevel),
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		IgnoreRelationshipsWhenMigrating:         false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		TranslateError:                           false,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})
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
	err = db.AutoMigrate(&models.KubernetesCluster{}, &models.KubernetesNamespace{}, &models.K8sResourceType{}, &models.K8sResourceConfig{}, &models.Permission{}, &models.Role{}, &models.RolePermission{}, &models.User{}, &models.UserRole{})
	if err != nil {
		return err
	}
	return nil
}
