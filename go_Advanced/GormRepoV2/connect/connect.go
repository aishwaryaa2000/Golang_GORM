package connect

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func SetupDB() *gorm.DB {

	dsn := "root:Pass@123@tcp(127.0.0.1:3306)/tsm5?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil
	} else {
		// fmt.Println("Connection successful")
		return db
	}

}

