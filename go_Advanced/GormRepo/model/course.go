package model

import uuid "github.com/satori/go.uuid"

type Course struct {
	CustomModel
	Name string
}

func NewCourse(name string) Course {
	id := uuid.NewV4()
	return Course{
		CustomModel: CustomModel{
			ID : id,
		},
		Name: name,
	}
}