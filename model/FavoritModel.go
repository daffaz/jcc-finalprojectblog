package model

import "gorm.io/gorm"

type Favorit struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	BlogID    int    `json:"blog_id"`
	Blog      Blog   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Commenter string `json:"commenter"`
	Content   string `json:"content" gorm:"type:text"`
}

func (b *Favorit) TableName() string {
	return "favorit"
}
