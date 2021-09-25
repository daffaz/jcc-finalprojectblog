package Model

import "gorm.io/gorm"

type Bantuan struct {
	gorm.Model
	UserID  int    `json:"user_id"`
	User    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content string `json:"content"`
}

func (b *Bantuan) TableName() string {
	return "bantuan"
}
