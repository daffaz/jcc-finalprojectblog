package Controllers

import (
	"jccblog/Model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllCategories responds with the list of all categories as JSON.
// GetAllCategories godoc
// @Summary Show all categories data.
// @Description This endpoint intended to show all categories data .
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {array} Model.Blog
// @Router /api/categories [get]
func GetAllCategories(c *gin.Context) {
	var category []Model.Category

	err := Model.GetCategory(&category)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, category)
	}
}

// CreateNewCategory godoc
// @Summary For creating a new category.
// @Description Use this endpoint with POST method if you want to create a new category Only admin can post category.
// @Tags category
// @Accept  json
// @Produce  json
// @Success 200 {object} Model.Category
// @Router /api/categories [post]
func CreateNewCategory(c *gin.Context) {
	var category Model.Category
	c.BindJSON(&category)

	err := Model.CreateCategory(&category)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, category)
	}
}
