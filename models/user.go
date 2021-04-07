package models

import (
	"fmt"
	"html"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"recipes-api.com/m/utils"
)

type User struct {
	gorm.Model
	Username 	string 		`gorm:"size:100;not null;unique" json:"username" binding:"required"`
	Email 		string 		`gorm:"size:100;not null;unique" json:"email" binding:"required"`
	Password 	string 		`gorm:"size:100;not null" binding:"required"`
	Recipes 	[]Recipe 	`gorm:"constraint:OnUpdate:CASCADE;foreignKey:UserID" json:"-"`
}

type UserReadModel struct {
	ID          uint
	Username 	string 	`gorm:"size:100;not null;unique" json:"username" binding:"required"`
	Email 		string 	`gorm:"size:100;not null;unique" json:"email" binding:"required"`
}

type UpdateUser struct {
	gorm.Model
	Username 	string 	`gorm:"size:100;not null" json:"username" binding:"required"`
	Email 		string 	`gorm:"size:100;not null" json:"email" binding:"required"`
	Password 	string 	`gorm:"size:100;not null" json:"password"`
}

type UpdateUserPassword struct {
	gorm.Model
	Password 	string 	`gorm:"size:100;not null" json:"password" binding:"required"`
}

type UserLoginAttempt struct {
	gorm.Model
	Email string 	`gorm:"size:100;not null;unique" json:"email" binding:"required"`
	Password string `gorm:"size:100;not null;unique" json:"password" binding:"required"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *User) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) ValidateEmail() (error) {

	emailValid, errMessage := utils.IsEmailValid(u.Email)

	if !emailValid {
		return fmt.Errorf("please provide a valid email. error message: %s", errMessage)
	}

	return nil
}

func (u *UpdateUser) ValidateEmail() (error) {

	emailValid, errMessage := utils.IsEmailValid(u.Email)

	if !emailValid {
		return fmt.Errorf("please provide a valid email. error message: %s", errMessage)
	}

	return nil
}

func (u *User) ValidateUsername() (error) {
	userNameValid, errMessage := utils.IsUserNameValid(u.Username)

	if !userNameValid {
		return fmt.Errorf("uername is invalid. error message: %s", errMessage)
	}

	return nil
}

func (u *User) Prepare() {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) BeforeSave(db *gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
func (u *UpdateUser) BeforeSave(db *gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (user *User) FindAllUsers(db *gorm.DB) (*[]UserReadModel, error) {

	var err error
	users:= []UserReadModel{}
	err = db.Debug().Model(&User{}).Find(&users).Error

	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *User) FindUserByID(db *gorm.DB, uid uint) (*UserReadModel, error) {
	var userReturnModel *UserReadModel
	err := db.Debug().Where("id = ?", uid).First(&userReturnModel).Error

	if err != nil {
		return nil, err
	}
	
	return userReturnModel, nil
}


func (u *User) UpdatePersonalDetails(db *gorm.DB, uid string, input UpdateUser) (*User, error) {

	err := u.BeforeSave(db)
	if err != nil {
		log.Fatal(err)
	}

	if err :=  db.Where("id = ?", uid).First(&u).Error; err != nil {
		return nil, err
	}

	db = db.Debug().Model(&u).UpdateColumns(
		map[string]interface{}{
			"email":     	input.Email,
			"username": 	input.Username,
		},
	)

	if db.Error != nil {
		return nil, db.Error
	}

	err = db.Debug().Model(&u).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) UpdateUserPassword(db *gorm.DB, uid string, newPassword string) (*User, error) {

	if err :=  db.Where("id = ?", uid).First(&u).Error; err != nil {
		return nil, err
	}

	u.Password = newPassword

	err := u.BeforeSave(db)
	if err != nil {
		log.Fatal(err)
	}


	db = db.Debug().Model(&u).UpdateColumn("password", u.Password)

	if db.Error != nil {
		return nil, db.Error
	}
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (user *User) EmailExists(db *gorm.DB, email string) (bool, error ){
	var userByEmailCount int64
	if err := db.Model(&user).Select("id").Where("email = ?", email).Count(&userByEmailCount).Error; err != nil {
       panic(err)
    } 

	if userByEmailCount > 0 {
		return true, nil
	}

	return false, nil
}

func (user *User) UserNameExists(db *gorm.DB, username string) (bool, error ){
	var userByUserNameCount int64
	if err := db.Model(&user).Select("id").Where("username = ?", username).Count(&userByUserNameCount).Error; err != nil {
        panic(err)
    } 

	if userByUserNameCount > 0 {
		return true, nil
	}

	return false, nil
}

