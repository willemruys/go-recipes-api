package models

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Title  		string 		`json:"title"`
	Ingredients string 		`json:"ingredients"`
	UserID		uint    	`gorm:"constraint:OnUpdate:CASCADE"`
	Likes		int 		`json:"likes"`
	Comments 	[]Comment 	`gorm:"constraint:OnUpdate:CASCADE;foreignKey:RecipeID"`
}


type CreateRecipe struct {
	gorm.Model
	Title  		string `json:"title" binding:"required"`
	Ingredients string `json:"ingredients" binding:"required"`
	UserID 		uint   `json:"userId"`
}

type UpdateRecipe struct {
	gorm.Model
	Title  		string `json:"title"`
	Ingredients string `json:"ingredients"`
}

func (r *Recipe) CreateRecipe(db *gorm.DB) (*Recipe, error)  {

	if err := db.Debug().Create(&r).Error; err != nil {
		return nil, err
	}

	db.Model(&r).Save(&r)

	return r, nil
}

func (r *Recipe) UpdateRecipe(db *gorm.DB, recipeId string, input UpdateRecipe) (*Recipe, error) {

	if err := db.Where("id = ?", recipeId).First(&r).Error; err != nil {
		return nil, err
	}

	if err := db.Debug().Model(&r).UpdateColumns(Recipe{Title: input.Title, Ingredients: input.Ingredients}).Error; err != nil {
		return nil, err
	}

	return r, nil
} 

func (r *Recipe) DeleteRecipe(db *gorm.DB, recipeId string) (error) {

	if err := db.Where("id = ?", recipeId).First(&r).Error; err != nil {
		return err
	}

	if err := db.Debug().Delete(&r).Where("ID = ?", recipeId).Error; err != nil {
		return err
	}

	return nil

}

func (r *Recipe) AddComment(db *gorm.DB, recipeId string, comment Comment) (*Recipe, error) {

	if err := db.Where("id = ?", recipeId).First(&r).Error; err != nil {
		return nil, err
	}

	comments := append(r.Comments, comment)

	r.Comments = comments

	if err := db.Save(&r).Error; err != nil {
		return nil, err
	}

	return r, nil

}

func (r *Recipe) GetRecipeComments(db *gorm.DB, recipeId string) ([]Comment, error) {

	if err :=  db.Where("id = ?", recipeId).First(&r).Error; err != nil {
		return nil, err
	}

	var comments []Comment

	if err := db.Debug().Model(&r).Association("Comments").Find(&comments); err != nil {
		return nil, err
	}

	return comments, nil
}