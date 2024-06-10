package models

import (
	"github.com/jinzhu/gorm"
)

// RolePermission 关联角色与权限的多对多关系
type RolePermission struct {
	gorm.Model
	RoleID       uint `gorm:"column:role_id;unique;not null;comment:角色ID"`       // 角色ID
	PermissionID uint `gorm:"column:permission_id;unique;not null;comment:权限ID"` // 权限ID
}
