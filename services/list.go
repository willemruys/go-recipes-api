package services

import "recipes-api.com/m/models"


func CreateList(list models.List) (models.List, error) {

	db := models.LoadDB()

	if err := db.Debug().Model(&list).Create(&list).Error; err != nil {
		return models.List{}, err
	}

	if err := db.Debug().Model(&list).Save(&list).Error; err != nil {
		return models.List{}, err
	}

	return list, nil
}

func GetList(listId string) (*models.List, error) {
	
	db := models.LoadDB()
	var list *models.List

	if err := db.Debug().Model(&list).Where("id = ?", listId).First(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
