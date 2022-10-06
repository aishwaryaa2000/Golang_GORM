package service

import (
	"gorm/app"
	"gorm/repository"
	"gorm.io/gorm"
)

type Service struct {
	gormRepo repository.Repository
	db       *gorm.DB
}

var singleInstanceOfSevice *Service

func getInstanceOfService() *Service {

	if singleInstanceOfSevice == nil {
		singleInstanceOfSevice = &Service{
			gormRepo: repository.NewRepository(),
			db:       app.GetDb(),
		}
	}

	return singleInstanceOfSevice

}

