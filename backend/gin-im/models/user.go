package models

import "gin-im/utils"

type User struct {
	Model
	UserName string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

// MakePassword 生成密码
func (u *User) MakePassword(password string) string {
	return utils.Encode256(password)
}

// ValidatePassword 验证密码
func (u *User) ValidatePassword(password string) bool {
	return utils.Encode256(password) == u.Password
}