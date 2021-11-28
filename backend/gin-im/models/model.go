package models

import "gin-im/utils"

type Model struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt utils.Time `json:"createdAt"`
	UpdatedAt utils.Time `json:"updatedAt"`
}
