package models

import (
	"gin-im/rbac/models"
	"gin-im/utils"
)

// User 用户表
type User struct {
	Model
	UserName    string        `gorm:"size:32;unique;not null" json:"username"`
	Password    string        `json:"password"`
	Status      uint          `gorm:"default:1;" json:"status"`
	IsSuperUser bool          `gorm:"default:false;" json:"is_super_user"`
	Roles       []models.Role `json:"roles" gorm:"many2many:user_roles"`
}

//func (User) TableName() string {
//	return "user"
//}

// MakePassword 生成密码
func (u *User) MakePassword(password string) string {
	crypto := utils.Crypto{}
	return crypto.Encode256(password)
}

// ValidatePassword 验证密码
func (u *User) ValidatePassword(password string) bool {
	crypto := utils.Crypto{}
	return crypto.Encode256(password) == u.Password
}
