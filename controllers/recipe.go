package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"recipes-api.com/m/auth"
	"recipes-api.com/m/models"
	"recipes-api.com/m/services"
)

func GetRecipes(c *gin.Context) {

	recipes, err := services.GetRecipes();

	if err != nil {
		c.JSON(http.StatusInternalServerError, "error retrieving recipes")
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": recipes})
}

func GetRecipe(c *gin.Context) {
	recipeId := c.Param("id")

	recipe, err := services.GetRecipe(recipeId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "error retrieving recipe")
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

func CreateRecipe(c *gin.Context) {

	input := models.CreateRecipe{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	userId, err := auth.ExtractTokenID(c);

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	recipeModel := models.Recipe{Title: input.Title, Ingredients: input.Ingredients, UserID: userId}

	recipe, err := services.CreateRecipe(recipeModel); 
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "recipe created", "recipe": recipe})

}

func UpdateRecipe(c *gin.Context) {

	recipeId := c.Param("id")

	var input models.UpdateRecipe

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe, err := services.GetRecipe(recipeId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}

	updatedRecipe, err := recipe.UpdateRecipe(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated": updatedRecipe})

}

func DeleteRecipe(c *gin.Context) {

	recipeId := c.Param("id")

	recipe, err := services.GetRecipe(recipeId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}
 
	if err := recipe.DeleteRecipe(); err != nil {
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
