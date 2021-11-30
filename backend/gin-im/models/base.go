package models

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ID uint `gorm:"primary_key" json:"id"`
}
