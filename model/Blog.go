package Model

import (
	"jccblog/Config"
	"log"
)

// GetBlogs... Fetch all blog data
func GetBlogs(blog *[]Blog) error {
	if err := Config.DB.Find(blog).Error; err != nil {
		return err
	}
	return nil
}

// CreateBlog... Create a new blog
func CreateBlog(blog *Blog) error {
	if err := Config.DB.Create(blog).Error; err != nil {
		return err
	}
	return nil
}

// GetBlogById... Fetch a single blog by it's blog_id
func GetBlogById(blog *Blog, id string) error {
	if err := Config.DB.Where("id = ?", id).First(blog).Error; err != nil {
		return err
	}
	return nil
}

// GetBlogBySlug... Fetch a single blog by it's slug
func GetBlogBySlug(blog *Blog, slug string) error {
	if err := Config.DB.Where("slug = ?", slug).First(blog).Error; err != nil {
		return err
	}
	return nil
}

// UpdateBlog... Update a single blog by it's id
func UpdateBlog(blog *Blog, id string) error {
	log.Println(blog)
	Config.DB.Save(blog)
	return nil
}

// DeleteBlog... Delete a single blog by it's id
func DeleteBlog(blog *Blog, id string) error {
	Config.DB.Where("id = ?", id).Delete(blog)
	return nil
}
