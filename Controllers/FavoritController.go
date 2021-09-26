package Controllers

import (
	"jccblog/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllFavoritByUID(c *gin.Context) {
	var favorit []Model.Favorit

	id := c.Params.ByName("id")
	err := Model.GetFavoritByUID(&favorit, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, favorit)
	}
}

func CreateNewFavorit(c *gin.Context) {
	var favorit Model.Favorit
	c.BindJSON(&favorit)

	err := Model.CreateFavorit(&favorit)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "favorit created", "data": favorit})
	}
}

func DeleteFavorit(c *gin.Context) {
	var favorit Model.Favorit
	id := c.Params.ByName("id")
	err := Model.DeleteFavorit(&favorit, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
