package controllers

import (
	"net/http"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindRecipes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var recipes []models.Recipe
	db.Find(&recipes)
	c.JSON(http.StatusOK, gin.H{"data": recipes})
}

func CreateRecipe(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.CreateRecipe
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recipe := models.Recipe{Title: input.Title, Ingredients: input.Ingredients}

	db.Create(&recipe)
	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

func UpdateRecipe(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var input models.UpdateRecipe


	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRecipe := models.Recipe{Title: input.Title, Ingredients: input.Ingredients}

	var recipeId = c.Param("id")

	var recipe models.Recipe

	if err := db.First(&recipe, recipeId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&recipe).Where("ID = ?", recipeId).Updates(updatedRecipe).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated": updatedRecipe})

}

func DeleteRecipe(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var recipeId string
	recipeId = c.Param("id")

	var recipe models.Recipe
	err := db.First(&recipe, recipeId).Error; if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 
	db.Delete(&recipe).Where("ID = ?", recipeId)

	c.JSON(http.StatusOK, gin.H{"Deleted":recipe })
}
