package db

import "gin-im/models"

func init() {
	//_ = DB.AutoMigrate(&models.User{})
	//_ = DB.AutoMigrate(&models.UserProfile{})
	
	DB.Migrator().CreateTable(&models.User{})
	DB.Migrator().CreateTable(&models.UserProfile{})
	DB.Migrator().RenameTable("users", "user")
	DB.Migrator().RenameTable("user_profiles", "user_profile")
}
