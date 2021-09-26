package Model

import (
	"jccblog/Config"
)

// GetCommentById... Get a comment
func GetCommentByBlogID(comment *[]Comment, id string) error {
	if err := Config.DB.Where("blog_id = ?", id).Find(comment).Error; err != nil {
		return err
	}
	return nil
}

func CreateComment(comment *Comment) error {
	if err := Config.DB.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

// DeleteComment... Delete a single comment by it's id
func DeleteComment(comment *Comment, id string) error {
	Config.DB.Where("id = ?", id).Delete(comment)
	return nil
}
