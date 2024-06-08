package models

import "github.com/jinzhu/gorm"

// User represents the User model with comments for each field.
type User struct {
	gorm.Model
	// 用户邮箱，唯一标识，不可为空
	Email string `gorm:"type:varchar(100);unique_index;not null;comment:邮箱" json:"email"`
	// 用户名，唯一，长度限制50字符，不可为空
	Username string `gorm:"type:varchar(50);unique_index;not null;comment:用户名" json:"username"`
	// 加密后的用户密码，不可为空
	Password string `gorm:"type:varchar(100);not null;comment:密码" json:"-"`
	// 用户名，长度限制50字符，不可为空
	FirstName string `gorm:"type:varchar(50);not null;comment:名" json:"first_name"`
	// 用户姓，长度限制50字符，不可为空
	LastName string `gorm:"type:varchar(50);not null;comment:姓" json:"last_name"`
}

// BeforeSave is a GORM callback, used for operations like hashing the password before saving.
func (u *User) BeforeSave(_ *gorm.DB) error {
	// 实际应用中在此处使用安全的密码哈希方法
	// u.Password = hashedPassword(u.Password)
	return nil
}
