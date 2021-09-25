package Controllers

import (
	"jccblog/Model"
	"jccblog/Utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func GetAllBlogs(c *gin.Context) {
	var blog []Model.Blog

	err := Model.GetBlogs(&blog)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, blog)
	}
}

func GetBlogDataById(c *gin.Context) {
	id := c.Params.ByName("id")
	var blog Model.Blog
	err := Model.GetBlogById(&blog, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusFound, blog)
	}
}

func GetBlogDataBySlug(c *gin.Context) {
	id := c.Params.ByName("slug")
	var blog Model.Blog
	err := Model.GetBlogBySlug(&blog, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusFound, blog)
	}
}

func CreateNewBlog(c *gin.Context) {
	var blog Model.Blog
	c.BindJSON(&blog)
	validatedPhotoBlog, err := Utils.ValidateUrl(blog.PhotoBlog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	blog.Slug = slug.Make(blog.Title)
	blog.PhotoBlog = validatedPhotoBlog
	log.Println(blog)
	err = Model.CreateBlog(&blog)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, blog)
	}
}

func UpdateUser(c *gin.Context) {
	var blog Model.Blog
	id := c.Params.ByName("id")
	err := Model.GetBlogById(&blog, id)
	if err != nil {
		c.JSON(http.StatusNotFound, blog)
	}
	c.BindJSON(&blog)
	err = Model.UpdateBlog(&blog, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, blog)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var blog Model.Blog
	id := c.Params.ByName("id")
	err := Model.DeleteBlog(&blog, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
