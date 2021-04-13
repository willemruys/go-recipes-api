package services

import (
	"recipes-api.com/m/models"
)

func GetRecipes() ([]models.Recipe, error) {

	db := models.LoadDB()
	var recipes []models.Recipe

	if err := db.Model(&recipes).Find(&recipes).Error; err != nil {	
		return nil, err
	}

	return recipes, nil
} 

func GetRecipe(recipeId string) (models.Recipe, error) {
	
	db := models.LoadDB()
	var recipe models.Recipe

	if err := db.Debug().Model(&recipe).Where("id = ?", recipeId).First(&recipe).Error; err != nil {
		return models.Recipe{}, err
	}

	return recipe, nil
}

func CreateRecipe(recipe models.Recipe) (models.Recipe, error) {
	
	db := models.LoadDB()

	if err := db.Debug().Model(&recipe).Create(&recipe).Error; err != nil {
		return models.Recipe{}, err
	}

	if err := db.Debug().Model(&recipe).Save(&recipe).Error;  err != nil {
		return models.Recipe{}, err
	}

	return recipe, nil
	
}