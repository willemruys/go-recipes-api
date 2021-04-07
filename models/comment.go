package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text      	string 			`gorm:"size:255;not null;unique" json:"text" binding:"required"`
	UserID 		uint
	UserName    string
	RecipeID  	uint
}

type UpdateComment struct {
	Text string `binding:"required"`
}

func (c *Comment) UpdateComment(updateComment UpdateComment) (*Comment, error) {

	db := LoadDB()
	if err := db.Debug().Model(&c).UpdateColumn("Text", updateComment.Text).Error; err != nil {
		return nil, err
	}

	return c, nil

}

func (c *Comment) DeleteComment() (error) {

	db := LoadDB()

	if err := db.Debug().Model(&c).Delete(&c).Error; err != nil {
		return err
	}

	return nil
}