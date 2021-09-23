package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	BlogID    Blog   `json:"blog_id"`
	Commenter string `json:"commenter"`
	Content   string `json:"content" gorm:"type:text"`
}

func (b *Comment) TableName() string {
	return "comment"
}
