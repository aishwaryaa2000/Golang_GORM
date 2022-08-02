package model

type User struct {
	CustomModel
	FName   string
	LName   string
	Courses []Course `gorm:"many2many:user_courses;"`
	Hobbies []Hobby  `gorm:"foreign_key:UserID"`
}
