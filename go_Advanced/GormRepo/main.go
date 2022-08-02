package main

import (
	
	// "gorm/connect"
	// "gorm/model"
	"gorm/service"
)

func main() {

	// createTableAndSetForiegnKey()
	service.GetAllUser()

	// service.AddUser()
	// service.GetUser()
	// service.GetAllUser()
	service.HardDeleteUser()
	
	// service.AddCourse()
	// service.GetAllCourses()
	// service.UpdateCourse()
	// service.DeleteCourse()
	service.GetAllCourses()

	// service.UpdateUser()
	 service.GetAllUser()
	//  service.DeleteUser()
	// service.GetAllUser()

}

// func createTableAndSetForiegnKey() {
// 	db := connect.SetupDB().Debug()
// 	db.CreateTable(&model.User{})
// 	db.CreateTable(&model.Course{})
// 	db.CreateTable(&model.Hobby{})
// 	db.Model(&model.Hobby{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
// }
