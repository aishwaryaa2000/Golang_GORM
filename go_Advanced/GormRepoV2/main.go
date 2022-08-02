package main

import (
	"gormv2/model"
	"gormv2/connect"
	"gormv2/service"

)

func main(){

	 db := connect.SetupDB().Debug()
	
	db.CreateTable(&model.User{})
	db.CreateTable(&model.Course{})
	db.CreateTable(&model.Hobby{})

	// service.AddUser()
}