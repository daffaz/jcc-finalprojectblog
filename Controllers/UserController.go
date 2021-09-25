package Controllers

import (
	"jccblog/Model"
	"jccblog/Utils"
	"log"
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
	log.Println(user)

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
		c.JSON(http.StatusCreated, gin.H{"message": "User created"})
	}
}
