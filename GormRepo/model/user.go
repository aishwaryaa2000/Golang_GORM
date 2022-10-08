package model

import (
	"gorm/encrypt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	CustomModel
	FName   string `json:"fname"`
	LName   string	`json:"lname"`
	Password string  `json:"password"`
	Courses []Course `gorm:"many2many:user_courses;constraint:OnDelete:CASCADE;" json:"courses"` 
	Hobbies []Hobby  `gorm:"foreign_key:UserID;constraint:OnDelete:CASCADE;" json:"hobbies"`
}

//hook to assign uuid to user id and hash user password
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	u.Password = encrypt.CreateHash(u.ID.String() + u.Password)
	return nil
}
