package models

import (
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	Title 		string  `json:"title" binding:"required"`
	Description string	`json:"description" binding:"required"`
	UserID 		uint64	`json:"userId"`	
	Recipes		[]Recipe `gorm:"many2many:list_recipes;"`
}

type UpdateList struct {
	gorm.Model
	Title  		string `json:"title"`
	Description string `json:"description"`
}


func (l *List) AddRecipeToList(recipe Recipe) (*List, error) {

	db := LoadDB()

	recipes := append(l.Recipes, recipe)

	l.Recipes = recipes

	if err := db.Debug().Save(&l).Error; err != nil {
		return nil, err
	}

	return l, nil
}

func (l *List) RemoveRecipeFromList(recipe Recipe) (*List, error) {

	db := LoadDB()

	if err := db.Debug().Model(&l).Association("Recipes").Delete(&recipe); err != nil {
		return nil, err
	}

	return l, nil

}

func (l *List) GetListRecipes() ([]Recipe, error) {

	db := LoadDB()

	var recipe []Recipe

	if err := db.Debug().Model(&l).Association("Recipes").Find(&recipe); err != nil {
		return nil, err
	}

	return recipe, nil
}

func (l *List) DeleteList() (error) {

	db := LoadDB()

	if err := db.Debug().Delete(&l).Error; err != nil {
		return err
	}

	return nil
}


func (l *List) UpdateList(input UpdateList) (*List, error) {
	
	db := LoadDB() 

	if err := db.Debug().Model(&l).UpdateColumns(List{Title: input.Title, Description: input.Description}).Error; err != nil {
		return nil, err
	}

	return l, nil

}