package service

import (
	"gorm/app"
	"gorm/repository"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	gormRepo repository.Repository
	db       *gorm.DB
}

var singleInstanceOfSevice *Service
var lock = &sync.Mutex{}

func getInstanceOfService() *Service {

	if singleInstanceOfSevice == nil {
		lock.Lock()
		/* 
		Mutex provides mutual exclusion to a resource, which means that only one goroutine can hold it at a time.
		Any goroutine must first acquire the mutex and then access the resource. 
		If a goroutine tries to acquire the mutex already held by another goroutine,
		it will be blocked and will wait for the mutex to be released.
		*/
        defer lock.Unlock()
		if singleInstanceOfSevice == nil {
			/* 
			check for nil singleInstanceOfSevice
			after the lock is acquired. This is to make sure that
			if more than one goroutine bypass the first check then only one goroutine
			is able to create the singleton instance otherwise each of the goroutine 
			will create its own instance of the single struct.
			*/
			appObject := app.NewWithEnv()
			singleInstanceOfSevice = &Service{
				gormRepo: repository.NewRepository(),
				db:       appObject.GetDb(),
			}
		}
	}

	return singleInstanceOfSevice

}

