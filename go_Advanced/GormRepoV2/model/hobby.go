package model

import(
	"gorm.io/gorm"
)

type Hobby struct{
  gorm.Model
  Number   string
  UserID  uint //This is the foriegn key
}