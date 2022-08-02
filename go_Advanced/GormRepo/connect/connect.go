package connect

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupDB() *gorm.DB {

	db, err := gorm.Open("mysql", "root:Pass@123@/tsm5?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return nil
	} else {
		// fmt.Println("Connection successful")
		return db
	}

}