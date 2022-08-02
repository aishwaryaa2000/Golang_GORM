package model
import (
	"github.com/satori/go.uuid"
)

type Hobby struct {
	Name   string
	UserID uuid.UUID `gorm:"foreign_key;not null;type:varchar(36)"` //Foreign key
	CustomModel
}

func NewHobby(name string) Hobby {
	id := uuid.NewV4()
	return Hobby{
		CustomModel: CustomModel{
			ID : id,
		},
		Name: name,
	}
}

/*

User HAS MANY hobbies so primary key ID should not exist for hobby - acc to class

User m2m courses 
courses ka fixed set hai so courses shoud have 

*/