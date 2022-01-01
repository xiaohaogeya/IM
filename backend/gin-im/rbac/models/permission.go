package models

// Permission 权限
type Permission struct {
	Model
	ApiUrl      string       `json:"api_url"`
	FrontUrl    string       `json:"front_url"`
	Method      string       `json:"method"`
	Title       string       `json:"title"`
	Rule        string       `json:"rule"`
	ParentId    uint         `json:"parentId"`
	Permissions []Permission `json:"-" gorm:"-"`
}
