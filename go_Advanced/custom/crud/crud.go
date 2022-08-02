package crud

import (
	"gorm/model"
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

func CreateTable(db *gorm.DB) {
	//Table having foriegn key should be dropped first
	db.DropTable(&model.User{})

	if !db.HasTable(&model.User{}) {
		db.AutoMigrate(&model.User{})
	}

}

func Insert(db *gorm.DB) {

	uid := uuid.NewV4()
	user := model.User{
		FirstName: "Aishwarya",
		LastName:  "Anand",
		Age:       18,
		model.CustomModel: model.CustomModel{
			ID: uid,
		},
	}
	db.Create(&user)

}
