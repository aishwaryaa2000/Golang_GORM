package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Course struct {
	CustomModel
	Name string `json:"cname" gorm:"uniqueIndex:idx_course_name;type:varchar(100)"`
}

func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.NewV4()
	return nil
}
