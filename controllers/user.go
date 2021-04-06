package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"recipes-api.com/m/models"
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

	if err := user.ValidateEmail(); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	if err := user.ValidateUsername(); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

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
		c.JSON(http.StatusNotAcceptable, gin.H{"response": "missing details"})
		return
	}

	if err := input.ValidateEmail(); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
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
	user := models.User{}
	userGotten, err := user.FindUserByID(db, uint32(uid))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userGotten})
}

func GetUserRecipes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": err})
		return
	}

	user := models.User{}

	recipes, err := user.UserRecipes(db, uid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "recipes": recipes})


}