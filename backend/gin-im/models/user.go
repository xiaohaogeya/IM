package models

import "gin-im/utils"

// User 用户表
type User struct {
	Model
	UserName    string `gorm:"size:32;unique;not null" json:"username"`
	Password    string `json:"password"`
	Status      uint   `gorm:"default:1;" json:"status"`
	IsSuperUser bool   `gorm:"default:false;" json:"is_super_user"`
}

// UserProfile 用户详情表
type UserProfile struct {
	Model
	UserId    uint   `json:"user_id"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Age       uint   `gorm:"default:0;" json:"age"`
	Mobile    string `gorm:"size:11;unique" json:"mobile"`
	Email     string `gorm:"size:32;unique" json:"email"`
	AvatarUrl string `json:"avatar_url"`
	NickName  string `gorm:"size:32;" json:"nick_name"`
}

// MakePassword 生成密码
func (u *User) MakePassword(password string) string {
	return utils.Encode256(password)
}

// ValidatePassword 验证密码
func (u *User) ValidatePassword(password string) bool {
	return utils.Encode256(password) == u.Password
}
