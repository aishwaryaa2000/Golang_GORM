package model
import (
	"time"
	"github.com/satori/go.uuid"

)

type CustomModel struct{
	ID        uuid.UUID `gorm:"primary_key,type:varchar(36"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

}

type User struct {
	CustomModel
	FirstName string
	LastName string
	Age int
}

