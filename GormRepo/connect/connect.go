package connect

import (
	"gorm.io/gorm"
  	"gorm.io/driver/mysql"
)

func SetupDB() *gorm.DB {

	dsn := "root:Pass@123@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
 	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db.Debug()

}