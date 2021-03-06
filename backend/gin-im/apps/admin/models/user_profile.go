package models

// UserProfile 用户详情表
type UserProfile struct {
	Model
	UserId    uint   `json:"user_id"`
	User      User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;polymorphicValue:master"`
	Age       uint   `gorm:"default:0;" json:"age"`
	Mobile    string `gorm:"size:11;" json:"mobile"`
	Email     string `gorm:"size:32;" json:"email"`
	AvatarUrl string `json:"avatar_url"`
	NickName  string `gorm:"size:32;" json:"nick_name"`
}

