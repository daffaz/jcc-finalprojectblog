package Controllers

import (
	"jccblog/Model"
	"jccblog/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers responds with the list of all users as JSON.
// GetAllUsers godoc
// @Summary Show all users, of course only user with admin role can see this
// @Description This endpoint intended to show all the registered users that is sent from the user. Only admin can see this
// @Tags blogs
// @Accept  json
// @Produce  json
// @Success 200 {array} Model.User
// @Router /api/users [get]
func GetAllUsers(c *gin.Context) {
	var user []Model.User

	err := Model.GetUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUserById responds with the data of user as JSON.
// GetUserById godoc
// @Summary Show user data by ID.
// @Description This endpoint intended to show user data that is sent from the user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} Model.User
// @Router /api/users/{id} [get]
func GetUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Model.User
	err := Model.GetUserById(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusFound, user)
	}
}

// CreateNewUser godoc
// @Summary For creating a new user (register).
// @Description Use this endpoint with POST method if you want to create a new user.
// @Tags blog
// @Accept  json
// @Produce  json
// @Success 200 {object} Model.User
// @Router /api/register [post]
func CreateNewUser(c *gin.Context) {
	var user Model.User
	c.BindJSON(&user)

	validatedPhotoProfile, err := Utils.ValidateUrl(user.PhotoProfile)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		user.PhotoProfile = validatedPhotoProfile
	}
	err = Model.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "User created", "data": user})
	}
}

// UpdateUser godoc
// @Summary For updating a user.
// @Description Use this endpoint with PUT method if you want to update a user data.
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} Model.User
// @Router /api/users/{id} [put]
func UpdateUser(c *gin.Context) {
	var user Model.User
	id := c.Params.ByName("id")
	err := Model.GetUserById(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Model.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser godoc
// @Summary For deleting a user.
// @Description Use this endpoint with DELETE method if you want to delete a user.
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} Model.User
// @Router /api/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	var user Model.User
	id := c.Params.ByName("id")
	err := Model.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
