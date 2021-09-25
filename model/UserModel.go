package Model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"username"`
	Password     string `json:"password"`
	Fullname     string `json:"fullname"`
	Role         string `json:"role" gorm:"default:writer"`
	PhotoProfile string `json:"photo_profile"`
	Blogs        []Blog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (b *User) TableName() string {
	return "user"
}
