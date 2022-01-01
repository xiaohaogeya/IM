package models

// Menu 菜单
type Menu struct {
	Model
	Title  string `gorm:"varchar(32)" json:"title"`
	Icon   string `json:"icon"`
	Weight int    `json:"weight"`
}
