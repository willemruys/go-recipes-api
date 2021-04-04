package models

import "gorm.io/gorm"

type User struct {
  gorm.Model
  ID     uint   `json:"id" gorm:"primary_key"`
  Email  string `json:"email"`
  Password string 
}