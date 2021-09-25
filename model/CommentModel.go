package Model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	BlogID    uint   `json:"blog_id"`
	Commenter string `json:"commenter"`
	Content   string `json:"content"`
}

func (b *Comment) TableName() string {
	return "comment"
}
