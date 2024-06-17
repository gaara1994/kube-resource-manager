package models

import (
	"github.com/jinzhu/gorm"
)

// User 代表用户模型
type User struct {
	BaseModel
	Username     string `json:"username" gorm:"column:user_name;not null;uniqueIndex;comment:用户名"` // 用户名，设置为唯一索引
	PasswordHash string `json:"password_hash" gorm:"column:password_hash;not null;comment:密码散列值"` // 密码散列值
	Email        string `json:"email" gorm:"column:email;comment:邮箱"`                                           // 邮箱，设置为唯一索引
}

// BeforeSave is a GORM callback, used for operations like hashing the password before saving.
func (u *User) BeforeSave(_ *gorm.DB) error {
	// 实际应用中在此处使用安全的密码哈希方法
	// u.Password = hashedPassword(u.Password)
	return nil
}

func (u *User) TableName() string {
	return "users"
}
