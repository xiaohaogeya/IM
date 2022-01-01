package db

import (
	models2 "gin-im/app/models"
)

func init() {
	_ = DB.AutoMigrate(&models2.User{})
	_ = DB.AutoMigrate(&models2.UserProfile{})

	//_ = DB.Migrator().CreateTable(&models.User{})
	//_ = DB.Migrator().CreateTable(&models.UserProfile{})
	//_ = DB.Migrator().RenameTable("users", "user")
	//_ = DB.Migrator().RenameTable("user_profiles", "user_profile")
}
