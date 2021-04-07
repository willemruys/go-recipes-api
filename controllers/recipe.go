package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"recipes-api.com/m/auth"
	"recipes-api.com/m/models"
)

func FindRecipes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var recipes []models.Recipe
	if err := db.Model(&recipes).Find(&recipes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": recipes})
}

func CreateRecipe(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input models.CreateRecipe

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId,err := auth.ExtractTokenID(c);

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}


	recipe := models.Recipe{Title: input.Title, Ingredients: input.Ingredients, UserID: userId}

	recipeRes, err := recipe.CreateRecipe(db);

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipeRes})
}

func UpdateRecipe(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	recipeId := c.Param("id")

	var input models.UpdateRecipe

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var recipe *models.Recipe
	recipe, err := recipe.UpdateRecipe(db, recipeId, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated": recipe})

}

func DeleteRecipe(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	recipeId := c.Param("id")

	var recipe models.Recipe
 
	if err := recipe.DeleteRecipe(db, recipeId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return		
	}

	c.JSON(http.StatusOK, gin.H{"Deleted": recipe})
}


func AddComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	recipeId := c.Param("id")
	userId, err := auth.ExtractTokenID(c);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input models.Comment
	var user models.User
	foundUser, err := user.FindUserByID(db, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = foundUser.ID
	input.UserName = foundUser.Username

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var recipe models.Recipe

	savedRecipe, err := recipe.AddComment(db, recipeId, input); 
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	

	c.JSON(http.StatusOK, gin.H{"message": savedRecipe})

}


func GetRecipeComments(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	recipeId := c.Param("id")

	recipe := models.Recipe{}

	comments, err := recipe.GetRecipeComments(db, recipeId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"recipe": recipe, "comments": comments})

}
