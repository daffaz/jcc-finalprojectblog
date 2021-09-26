package Controllers

import (
	"jccblog/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllFavoritByUID responds with the list of all favorit as JSON By ID.
// GetAllFavoritByUID godoc
// @Summary Show blogs favorit data by User ID.
// @Description This endpoint intended to show what user favorite's blog
// @Tags favorit
// @Accept  json
// @Produce  json
// @Param id path int true "Favorit ID"
// @Success 200 {array} Model.Favorit
// @Router /api/favorit/user_id/{id} [get]
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

// CreateNewFavorit godoc
// @Summary For creating a new favorit.
// @Description Use this endpoint with POST method if you want to create a new favorit.
// @Tags blog
// @Accept  json
// @Produce  json
// @Success 200 {object} Model.Favorit
// @Router /api/favorit [post]
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

// DeleteFavorit godoc
// @Summary For deleting a favorit.
// @Description Use this endpoint with DELETE method if you want to delete a favorit.
// @Tags favorit
// @Produce  json
// @Param id path int true "Favorit ID"
// @Success 200 {object} Model.Favorit
// @Router /api/favorit/{id} [delete]
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
