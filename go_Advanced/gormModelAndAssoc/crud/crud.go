package crud

import (
	"gorm/model"
	"github.com/jinzhu/gorm"
)

func CreateTable(db *gorm.DB) {
	//Table having foriegn key should be dropped first
	// db.DropTable(&model.Profile{})
	// db.DropTable(&model.User{})

	// if !db.HasTable(&model.User{}) {
	// 	db.AutoMigrate(&model.User{})
	// }

	// if !db.HasTable(&model.Profile{}) {
	// 	db.CreateTable(&model.Profile{})
	// }

	db.CreateTable(&model.User{})
	db.CreateTable(&model.Address{})
	db.CreateTable(&model.CreditCard{})
	
}

func MakeForeignKey(db *gorm.DB) {

	db.Model(&model.Address{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&model.CreditCard{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}

func Insert(db *gorm.DB) {

	tempUser := &model.User{
		CreditCard: model.CreditCard{
			Number: "123121864",
		},
		FirstName: "Aishwarya",
		LastName:  "Anand",
		Age:       22,
		Address: []model.Address{
			{
				City:  "Seawoods",
				State: "Mumbai",
			},
			{
				City:  "Mumbai",
				State: "Maharastra",
			},
		},
	}

	db.Create(tempUser)
	tempUser = &model.User{
		CreditCard: model.CreditCard{
			Number: "123121864",
		},
		FirstName: "Aayush",
		LastName:  "Anand",
		Age:       22,
		Address: []model.Address{
			{
				City:  "jbp",
				State: "Mp",
			},
			{
				City:  "Mumbai",
				State: "Maharastra",
			},
		},
	}

	db.Create(tempUser)
}

