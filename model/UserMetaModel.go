package Model

import "gorm.io/gorm"

type UserMeta struct {
	gorm.Model
	UserID       int    `json:"email"`
	User         User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Country      string `json:"country"`
	Gender       string `json:"gender"`
	PhotoProfile string `json:"photo_profile"`
}

func (b *UserMeta) TableName() string {
	return "user_meta"
}
