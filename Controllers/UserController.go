package Controllers

import (
	"jccblog/Model"
	"jccblog/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var user []Model.User

	err := Model.GetUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

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

//DeleteUser ... Delete the user
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
