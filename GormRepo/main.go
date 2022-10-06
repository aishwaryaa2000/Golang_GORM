package main

import (

	"gorm/connect"
	"gorm/model"
	"fmt"
	"gorm/router"
	"net/http"
)

func init() {
	db := connect.SetupDB().Debug()

	if !db.Migrator().HasTable(&model.User{}){
		db.AutoMigrate(&model.User{})
		fmt.Println("User table created")
	}
	if !db.Migrator().HasTable(&model.Course{}){
		db.AutoMigrate(&model.Course{})
		fmt.Println("Course table created")

	}
	if !db.Migrator().HasTable(&model.Hobby{}){
		db.AutoMigrate(&model.Hobby{})
		fmt.Println("Hobby table created")

	}
}

func main() {

	handler := router.CreateRoute()

	srv := &http.Server{
        Addr:         "0.0.0.0:8080",
        Handler: handler, // Pass our instance of gorilla/mux in.
    }

	
	fmt.Println("Starting server at 8080")

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
	

}

