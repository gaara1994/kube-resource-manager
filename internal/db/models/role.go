package models

import (
	"github.com/jinzhu/gorm"
)

// Role 代表角色模型
type Role struct {
	gorm.Model
	Name        string `gorm:"column:name;unique;not null;comment:角色名称"` // 角色名称，唯一且非空
	Description string `gorm:"column:description;comment:角色描述"`          // 角色描述
}
