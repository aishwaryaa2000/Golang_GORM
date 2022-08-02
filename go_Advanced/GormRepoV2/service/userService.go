package service

import (
	"fmt"
	"gormv2/connect"
	"gormv2/model"
	"gormv2/repository"
)

func getRepo() repository.Repository{
repo := repository.NewRepository() 
return repo
}
func AddUser(){
	db:= connect.SetupDB()
	uow := repository.NewUnitOfWork(db,false)
	defer uow.Complete()

	userr := model.User{
		FName: "Aayush",
		LName: "Kutty",
		ID: 2,
		Courses : []model.Course{
			model.Course{
			ID : 11,
			Name : "maths",
			},
			model.Course{
			ID : 22,
			Name : "hindi",
			},
		},
		
	}
	repo := getRepo()
	err := repo.Add(uow,userr)
	if err!=nil{
		fmt.Println(err)
	}
	
	uow.Commit()


}

/*

user ke service ke through hobbies

*/