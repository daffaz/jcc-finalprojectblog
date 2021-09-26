package Model

import (
	"gorm.io/gorm"
)

// Blog represents the model for a blog
type Blog struct {
	gorm.Model
	UserID     uint      `json:"user_id"`
	User       User      `json:"-"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	PhotoBlog  string    `json:"photo_blog"`
	Content    string    `json:"content"`
	Comments   []Comment `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CategoryID int       `json:"category_id"`
	Category   Category  `json:"-"`
}

func (b *Blog) TableName() string {
	return "blog"
}
