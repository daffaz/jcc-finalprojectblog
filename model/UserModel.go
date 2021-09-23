package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `json:"email"`
	Password     string `json:"assword"`
	Fullname     string `json:"fullname"`
	Role         string `gorm:"default:writer"`
	PhotoProfile string `json:"photo_profile"`
	Blogs        []Blog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (b *User) TableName() string {
	return "user"
}
