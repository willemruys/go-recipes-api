package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"recipes-api.com/m/auth"
	"recipes-api.com/m/models"
	"recipes-api.com/m/services"
)

func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailExists, err := user.EmailExists(user.Email)

	if emailExists || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email adres already exists"})
		return
	}

	userNameExists, err := user.UserNameExists(user.Username)

	if userNameExists || err != nil {
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

	userId, err := auth.ExtractTokenIDFromGinContext(c)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"response": err.Error()})
		return
	}

	var input models.UpdateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"response": "missing details"})
		return
	}

	if err := input.ValidateEmail(); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	user, err := services.GetUser(userId)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	emailExists, err := user.EmailExists(input.Email)

	if emailExists {
		c.JSON(http.StatusNotAcceptable, gin.H{"response": "email exists"})
		return
	}
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"response": err.Error()})
		return
	}
	
	if err := user.UpdatePersonalDetails(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": user})
}

func UpdatePassword(c *gin.Context) {

	userId, err := auth.ExtractTokenIDFromGinContext(c)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"response": err.Error()})
		return
	}

	var input models.UpdateUserPassword
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"response": err.Error()})
		return
	}

	user, err := services.GetUser(userId)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	updatedUser, err := user.UpdateUserPassword(userId, input.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": updatedUser})
}


func GetUser(c *gin.Context) {

	userId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"response": err.Error()})
		return
	}

	user, err := services.GetUser(userId)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUsers(c *gin.Context) {

	users, err := services.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserRecipes(c *gin.Context) {

	userId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"response": err.Error()})
		return
	}
	user, err := services.GetUser(userId)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	recipes, err := user.UserRecipes()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user.ID, "recipes": recipes})

}