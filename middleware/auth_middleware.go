package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"recipes-api.com/m/auth"
	"recipes-api.com/m/models"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "no valid jwt"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func RecipeOwnershipAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		recipeId := c.Param("id")

		var recipeModel models.Recipe

		db.Model(recipeModel).Where("id = ?", recipeId).First(&recipeModel);

		userId, err := auth.ExtractTokenID(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "Error retrieving JWT token"})
			c.Abort()
			return
		}

		if userId != recipeModel.UserID {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "You do not have permission to edit this recipe"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func OwnProfileOwnerShip() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		userId := c.Param("id")

		var user models.User

		db.Model(user).Where("id = ?", userId).First(&user);

		userIdFromJwt, err := auth.ExtractTokenID(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "Error retrieving JWT token"})
			c.Abort()
			return
		}

		if userIdFromJwt != user.ID {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "You do not have permission to edit this user"})
			c.Abort()
			return
		}

		c.Next()
	}
}
