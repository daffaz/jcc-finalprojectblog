package Controllers

import (
	"jccblog/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCommentByBlogID responds with the list of all comment as JSON.
// GetCommentByBlogID godoc
// @Summary Show comment of blogs data by ID.
// @Description This endpoint intended to show comments of some blogs data that is sent from the user
// @Tags comment
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Success 200 {object} Model.Comment
// @Router /api/comment/{id} [get]
func GetCommentByBlogID(c *gin.Context) {
	id := c.Params.ByName("id")
	var comment []Model.Comment
	err := Model.GetCommentByBlogID(&comment, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusFound, comment)
	}
}

// CreateNewComment godoc
// @Summary For creating a new blog's comment.
// @Description Use this endpoint with POST method if you want to create a new comment.
// @Tags blog
// @Accept  json
// @Produce  json
// @Success 200 {object} Model.Comment
// @Router /api/comment [post]
func CreateNewComment(c *gin.Context) {
	var comment Model.Comment
	c.BindJSON(&comment)

	err := Model.CreateComment(&comment)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, comment)
	}
}

// DeleteComment godoc
// @Summary For deleting a blog.
// @Description Use this endpoint with DELETE method if you want to delte a  comment.
// @Tags comment
// @Produce  json
// @Param id path int true "Comment ID"
// @Success 200 {object} Model.Comment
// @Router /api/comment/{id} [delete]
func DeleteComment(c *gin.Context) {
	var comment Model.Comment
	id := c.Params.ByName("id")
	err := Model.DeleteComment(&comment, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
