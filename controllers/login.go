package controllers

import (
	"net/http"

	"example.com/m/auth"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userLoginAttempt := models.UserLoginAttempt{}

	if err := c.ShouldBindJSON(&userLoginAttempt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error

	user := models.User{}

	err = db.Model(&user).Where("email = ?", userLoginAttempt.Email).Take(&user).Error; if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"response": "user does not exist"})
		return
	}

	user.Prepare()

	token, err := auth.SignIn(user, userLoginAttempt.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "failed login"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
