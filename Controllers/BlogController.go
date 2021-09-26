package Controllers

import (
	"jccblog/Model"
	"jccblog/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

// GetAllBlogs responds with the list of all blogs as JSON.
// GetAllBlogs godoc
// @Summary Show all blogs data sent from the user.
// @Description This endpoint intended to show all the blogs data that is sent from the user. Only registered user can see what other already post
// @Tags blogs
// @Accept  json
// @Produce  json
// @Success 200 {array} Model.Blog
// @Router /api/blogs [get]
func GetAllBlogs(c *gin.Context) {
	var blog []Model.Blog

	err := Model.GetBlogs(&blog)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, blog)
	}
}

// GetBlogDataById responds with the blog data as JSON.
// GetBlogDataById godoc
// @Summary Show blog data by ID.
// @Description This endpoint intended to show blog data that is sent from the user
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Success 200 {object} Model.Blog
// @Router /api/blogs/{id} [get]
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

// GetBlogDataBySlug responds with the list of all blogs as JSON.
// GetBlogDataBySlug godoc
// @Summary Show blog data by slug.
// @Description This endpoint intended to show blog data that is sent from the user
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param slug path string true "Blog slug"
// @Success 200 {object} Model.Blog
// @Router /api/blogs/{slug} [get]
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

// CreateNewBlog godoc
// @Summary For creating a new blog.
// @Description Use this endpoint with POST method if you want to create a new blog.
// @Tags blog
// @Accept  json
// @Produce  json
// @Success 200 {object} Model.Blog
// @Router /api/blogs [post]
func CreateNewBlog(c *gin.Context) {
	var blog Model.Blog
	c.BindJSON(&blog)
	validatedPhotoBlog, err := Utils.ValidateUrl(blog.PhotoBlog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	blog.Slug = slug.Make(blog.Title)
	blog.PhotoBlog = validatedPhotoBlog
	err = Model.CreateBlog(&blog)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, blog)
	}
}

// UpdateBlog godoc
// @Summary For updating a blog.
// @Description Use this endpoint with PUT method if you want to update a  blog.
// @Tags blog
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Success 200 {object} Model.Blog
// @Router /api/blogs/{id} [put]
func UpdateBlog(c *gin.Context) {
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

// DeleteBlog godoc
// @Summary For deleting a blog.
// @Description Use this endpoint with DELETE method if you want to delete a  blog.
// @Tags blog
// @Produce  json
// @Param id path int true "Blog ID"
// @Success 200 {object} Model.Blog
// @Router /api/blogs/{id} [delete]
func DeleteBlog(c *gin.Context) {
	var blog Model.Blog
	id := c.Params.ByName("id")
	err := Model.DeleteBlog(&blog, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
