package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Title  		string 	`json:"title"`
	Ingredients string 	`json:"ingredients"`
}

type CreateRecipe struct {
	gorm.Model
	Title  		string `json:"title" binding:"required"`
	Ingredients string `json:"ingredients" binding:"required"`
}

type UpdateRecipe struct {
	gorm.Model
	Title  		string `json:"title"`
	Ingredients string `json:"ingredients"`
}