package models

import (
	"gin-im/utils"
)

type Model struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt utils.Time `json:"created_at"`
	UpdatedAt utils.Time `json:"updated_at"`
}
