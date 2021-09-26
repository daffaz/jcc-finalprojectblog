package Model

import "gorm.io/gorm"

// Favorit represents the model for a favorit
type Favorit struct {
	gorm.Model
	UserID uint `json:"user_id"`
	BlogID int  `json:"blog_id"`
	Blog   Blog `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (b *Favorit) TableName() string {
	return "favorit"
}
