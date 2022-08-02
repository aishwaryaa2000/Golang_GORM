package model	

import (
	"github.com/satori/go.uuid"
	"time"
)

type CustomModel struct{	
	ID        uuid.UUID `gorm:"primary_key; type:varchar(36); not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index"`
}