package services

import "recipes-api.com/m/models"


func GetComment(commentId string) (*models.Comment, error) {

	db := models.LoadDB()

	var comment *models.Comment
	if err := db.Debug().Model(&comment).Where("id = ?", commentId).First(&comment).Error; err != nil {
		return nil, err
	}

	return comment, nil

}