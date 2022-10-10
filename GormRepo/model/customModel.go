package model	

import (
	"github.com/satori/go.uuid"
	"time"
	"gorm.io/gorm"
)

type CustomModel struct{	
	ID        uuid.UUID `gorm:"primary_key; type:varchar(36); not null" json:"id"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `sql:"index"`
}