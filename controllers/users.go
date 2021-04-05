package controllers

import (
	"net/http"
	"strconv"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser (c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Prepare()

	userCreated, err := user.SaveUser(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userCreated})
}


func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": err})
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByID(db, uint32(uid))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userGotten})
}