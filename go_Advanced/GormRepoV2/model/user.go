package model
import(
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	FName string
	LName string
	Courses []Course `gorm:"many2many:user_courses;"`
	Hobbies []Hobby

}