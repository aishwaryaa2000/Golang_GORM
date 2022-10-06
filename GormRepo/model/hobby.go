package model
import (
	"github.com/satori/go.uuid"
)

type Hobby struct {
	Id int64 `gorm:"autoIncrement:true"`
	Name   string   `gorm:"uniqueIndex:idx_name;type:varchar(100)" json:"hname"`
	UserID uuid.UUID `gorm:"foreign_key;not null;uniqueIndex:idx_name;type:varchar(100)"` //Foreign key
	//we make unique index for userid+hobbyname coz aishwarya-dance and aayush-dance should be allowed
}

/*
User HAS MANY hobbies so primary key ID should not exist for hobby - acc to class
So here ID is just for logging purposes

User m2m courses 
courses ka fixed set hai so courses shoud have ID

*/