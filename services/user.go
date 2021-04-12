package services

import "recipes-api.com/m/models"

func GetUser(userId uint64) (*models.User, error) {
	db := models.LoadDB()

	var user *models.User

	if err := db.Debug().Model(&user).Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUsers() (*[]models.User, error) {
	db := models.LoadDB()

	var users *[]models.User

	if err := db.Debug().Model(&users).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}