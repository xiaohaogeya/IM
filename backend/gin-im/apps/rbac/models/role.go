package models

// Role 角色
type Role struct {
	Model
	Title       string       `json:"title"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"`
}
