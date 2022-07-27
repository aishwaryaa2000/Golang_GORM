package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
  
  func main() {
	
	db, err := gorm.Open("mysql", "root:Pass@123@/db?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("Connection successful")
	}
	defer db.Close()
}