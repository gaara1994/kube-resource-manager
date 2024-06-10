package models

import (
	"github.com/jinzhu/gorm"
)

// Permission 代表权限模型
type Permission struct {
	gorm.Model
	Name        string `gorm:"column:name;unique;not null;comment:权限名称"` // 权限名称，唯一且非空
	Description string `gorm:"column:description;comment:权限描述"`          // 权限描述
}
