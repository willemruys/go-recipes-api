package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text      	string 			`gorm:"size:255;not null;unique" json:"text" binding:"required"`
	UserID 		uint
	UserName    string
	RecipeID  	uint
}