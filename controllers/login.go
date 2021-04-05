package controllers

import (
	"net/http"

	"example.com/m/auth"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Prepare()

	token, err := auth.SignIn(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": err})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
