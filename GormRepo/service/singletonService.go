package service

import (
	"gorm/connect"
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
			db:       connect.SetupDB().Debug(),
		}
	}

	return singleInstanceOfSevice
	
}

/*
Singleton design pattern
Didn't add Mutex part as go routines are not included for now
*/