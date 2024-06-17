package models

// Permission 代表权限模型
type Permission struct {
	BaseModel
	Name        string `json:"name" gorm:"column:name;unique;not null;comment:权限名称"` // 权限名称，唯一且非空
	Description string `json:"description" gorm:"column:description;comment:权限描述"`          // 权限描述
}
