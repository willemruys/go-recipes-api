package models

import "gorm.io/gorm"

type Recipe struct {
  gorm.Model
  ID     uint   `json:"id" gorm:"primary_key"`
  Title  string `json:"title"`
  Ingredients string `json:"ingredients"`
}

type CreateRecipe struct {
	Title  string `json:"title" binding:"required"`
	Ingredients string `json:"ingredients" binding:"required"`
}

type UpdateRecipe struct {
	Title string `json:"title"`
	Ingredients string `json:"ingredients"`
}