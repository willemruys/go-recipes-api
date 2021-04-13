package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"recipes-api.com/m/auth"
	"recipes-api.com/m/models"
	"recipes-api.com/m/services"
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
		recipeId := c.Param("id")
		
		recipe, err := services.GetRecipe(recipeId)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"response": "Recipe not found"})
			c.Abort()
			return
		}

		userId, err := auth.ExtractTokenIDFromGinContext(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "Error retrieving JWT token"})
			c.Abort()
			return
		}

		if userId != recipe.UserID {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "You do not have permission to edit this recipe"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func CommentOwnerShip() gin.HandlerFunc {
	return func(c *gin.Context) {

		commentId := c.Param("id")

		comment, err := services.GetComment(commentId); 
		
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"response": "Error retrieving comment"})
			c.Abort()
			return
		}

		userId, err := auth.ExtractTokenIDFromGinContext(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "Error retrieving JWT token"})
			c.Abort()
			return
		}

		if userId != comment.UserID {
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

		userIdFromJwt, err := auth.ExtractTokenIDFromGinContext(c)

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

func ListOwnerShip() gin.HandlerFunc {
	return func(c *gin.Context) {

		listId := c.Param("id")

		list, err := services.GetList(listId); 
		
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"response": "Error retrieving list"})
			c.Abort()
			return
		}

		userId, err := auth.ExtractTokenIDFromGinContext(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "Error retrieving JWT token"})
			c.Abort()
			return
		}

		if userId != list.UserID {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "You do not have permission to edit this list"})
			c.Abort()
			return
		}

		c.Next()
	}
}


