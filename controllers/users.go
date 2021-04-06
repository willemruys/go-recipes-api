package controllers

import (
	"net/http"
	"strconv"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailExists, err := user.EmailExists(db, user.Email)

	if emailExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email adres already exists"})
		return
	}

	userNameExists, err := user.UserNameExists(db, user.Username)

	if userNameExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This username adres already exists"})
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

func UpdateUserPersonalDetails(c *gin.Context) {

	userId := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	var input models.UpdateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"response": err.Error()})
		return
	}

	user := models.User{}

	emailExists, err := user.EmailExists(db, input.Email)
	if emailExists {
		c.JSON(http.StatusNotAcceptable, gin.H{"response": "email exists"})
		return
	}
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"response": err.Error()})
		return
	}

	updatedUser, err := user.UpdatePersonalDetails(db, userId, input)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": updatedUser})
}

func UpdatePassword(c *gin.Context) {

	userId := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	var input models.UpdateUserPassword
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"response": err.Error()})
		return
	}

	user := models.User{}

	updatedUser, err := user.UpdateUserPassword(db, userId, input.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": updatedUser})
}


func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": err})
		return
	}
	user := models.UserReadModel{}
	userGotten, err := user.FindUserByID(db, uint32(uid))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userGotten})
}