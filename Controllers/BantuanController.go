package Controllers

import (
	"jccblog/Model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllBantuan responds with the list of all albums as JSON.
// GetAllBantuan godoc
// @Summary Show all 'bantuan' data sent from the user.
// @Description This endpoint intended to show all the 'bantuan' data that is sent from the user.
// @Tags bantuan
// @Accept  json
// @Produce  json
// @Success 200 {array} Model.Bantuan
// @Router /api/bantuan [get]
func GetAllBantuan(c *gin.Context) {
	var bantuan []Model.Bantuan

	err := Model.GetBantuan(&bantuan)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, bantuan)
	}
}

// CreateNewBantuan godoc
// @Summary For creating a new bantuan.
// @Description Use this endpoint with POST method if you want to create a new bantuan.
// @Tags bantuan
// @Accept  json
// @Produce  json
// @Success 200 {object} Model.Bantuan
// @Router /api/bantuan [post]
func CreateNewbantuan(c *gin.Context) {
	var bantuan Model.Bantuan
	c.BindJSON(&bantuan)

	err := Model.CreateBantuan(&bantuan)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, bantuan)
	}
}
