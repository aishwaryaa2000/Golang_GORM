package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Course struct {
	CustomModel
	Name string `json:"cname"`
}

func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.NewV4()
	return nil
}
