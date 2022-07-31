package model

import(
	"github.com/jinzhu/gorm"
)
// type User struct {
// 	gorm.Model
// 	Name string
//   }
  
//   type Profile struct {
// 	gorm.Model
// 	Name      string
// 	User      User `gorm:"foreignkey:UserRefer"` // use UserRefer as foreign key
// 	UserRefer uint
//   }

type User struct {
  gorm.Model
  FirstName string
  LastName  string
  Age       int
  //Has Many With Address
  Address []Address
  //Has One Relation with Credit Card
  CreditCard CreditCard
}

type CreditCard struct {
  Number string
  UserId int `gorm:"foreign_key;not null "`
}

type Address struct {
  UserId int `gorm:"foreign_key; not null"`
  City   string
  State  string
}


  //If we are using primary key of gorm.model then we cannot set another field as primary key
  /*
	userId is automatically taken as foriegn key 
	Types should be compatible
	In gorm.model, ID is not null and unsigned data type so foreign keys should have same types in db
 */ 