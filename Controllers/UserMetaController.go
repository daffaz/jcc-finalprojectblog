package Controllers

import (
	"jccblog/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserMeta responds with the data of users meta as JSON.
// GetUserMeta godoc
// @Summary Show usermeta data .
// @Description This endpoint intended to show usermeta from users.
// @Tags usermeta
// @Accept  json
// @Produce  json
// @Param id path int true "UserMetaID"
// @Success 200 {object} Model.UserMeta
// @Router /api/usermeta [get]
func GetUserMeta(c *gin.Context) {
	var user Model.UserMeta

	id := c.Params.ByName("id")
	err := Model.GetUserMetaById(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateNewUserMeta godoc
// @Summary For creating a new user meta data.
// @Description Use this endpoint with POST method if you want to create a new user meta.
// @Tags usermeta
// @Accept  json
// @Produce  json
// @Success 200 {object} Model.UserMeta
// @Router /api/usermeta [post]
func CreateNewUserMeta(c *gin.Context) {
	var user Model.UserMeta
	c.BindJSON(&user)

	err := Model.CreateUserMeta(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "User meta created", "data": user})
	}
}
