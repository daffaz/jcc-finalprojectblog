package Model

import "gorm.io/gorm"

// UserMeta represents the model for a user meta
type UserMeta struct {
	gorm.Model
	UserID      int    `json:"user_id"`
	User        User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Country     string `json:"country"`
	Gender      string `json:"gender"`
	PhotoBanner string `json:"photo_banner"`
}

func (b *UserMeta) TableName() string {
	return "user_meta"
}
