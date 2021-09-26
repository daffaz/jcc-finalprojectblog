package Controllers

import (
	"jccblog/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

//DeleteComment ... Delete the comment
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
