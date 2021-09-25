package Controllers

import (
	"jccblog/Model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	var category []Model.Category

	err := Model.GetCategory(&category)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, category)
	}
}

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
