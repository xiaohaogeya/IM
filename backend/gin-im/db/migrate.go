package db

import (
	adminModels "gin-im/apps/admin/models"
	rbacModels "gin-im/apps/rbac/models"
)

func init() {
	_ = DB.AutoMigrate(&adminModels.User{})
	_ = DB.AutoMigrate(&adminModels.UserProfile{})

	_ = DB.AutoMigrate(&rbacModels.Menu{})
	_ = DB.AutoMigrate(&rbacModels.Permission{})
	_ = DB.AutoMigrate(&rbacModels.Role{})
	//_ = DB.AutoMigrate(&rbacModels.RolePermission{})

	//_ = DB.Migrator().CreateTable(&models.User{})
	//_ = DB.Migrator().CreateTable(&models.UserProfile{})
	//_ = DB.Migrator().RenameTable("users", "user")
	//_ = DB.Migrator().RenameTable("user_profiles", "user_profile")
}
