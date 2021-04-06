package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"recipes-api.com/m/auth"
	"recipes-api.com/m/models"
)

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userLoginAttempt := models.UserLoginAttempt{}

	if err := c.ShouldBindJSON(&userLoginAttempt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	if err := db.Model(&user).Where("email = ?", userLoginAttempt.Email).Take(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"response": "user does not exist"})
		return
	}

	user.Prepare()

	token, err := auth.AssignToken(user, userLoginAttempt.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "failed login"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
