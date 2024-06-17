package models

// UserRole 关联用户与角色的多对多关系
type UserRole struct {
	BaseModel
	UserID uint `gorm:"column:user_id;not null;comment:用户ID"` // 用户ID
	RoleID uint `gorm:"column:role_id;not null;comment:角色ID"` // 角色ID
}
