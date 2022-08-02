package main

import (
  "gorm/connect"
  "github.com/jinzhu/gorm"
)

type Userr struct {
	gorm.Model
	Name string
  }
  
  // `Profile` belongs to `User`, `UserID` is the foreign key
  type Profilee struct {
	gorm.Model
	UserID uint
	User   Userr
	Name   string
  }

  func MakeForeignKey(db *gorm.DB) {

	db.Model(&Profilee{}).AddForeignKey("user_id", "userrs(id)", "RESTRICT", "RESTRICT")
}

  func main(){
	db := connect.SetupDB()
	//db.CreateTable(Profilee{})
	//db.CreateTable(Userr{})

	MakeForeignKey(db)

	tempUser := Profilee{
		Name : "AishwaryaProfile",
		User : Userr{
			Name : "username",
		},
	
	}

	db.Create(&tempUser)


}