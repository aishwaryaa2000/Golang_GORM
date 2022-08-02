package service

import (
	"fmt"
	"gorm/model"
	"gorm/repository"
	"github.com/satori/go.uuid"
)


var serviceInstanceCourse = getInstanceOfService()

func AddCourse(){

	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, false)
	defer uow.Complete()
	course1 := model.NewCourse("Geography")
	// course2 := model.NewCourse("Politcs")
		
	err := serviceInstanceCourse.gormRepo.Add(uow, &course1)
	if err != nil {
		fmt.Println(err)
	}

	uow.Commit()
}


func GetCourse() {

	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, true)
	defer uow.Complete()

	var currCourse model.Course
	var preloadAssoc = []string{}
	currID ,_ := uuid.FromString("a5903424-cd3e-4025-abdb-e40084ddc7b9")
	err := serviceInstanceCourse.gormRepo.Get(uow, &currCourse, currID, preloadAssoc)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ID is : ", currCourse.ID)
	fmt.Println("Name is : ", currCourse.Name)

	uow.Commit()

}

func GetAllCourses() {

	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, true)
	defer uow.Complete()

	var courses []model.Course
	var preloadAssoc = []string{}
	err := serviceInstanceCourse.gormRepo.GetAll(uow, &courses, preloadAssoc)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nLIST OF ALL COURSES")
	for _, currCourse := range courses {
		fmt.Println("\nID is : ", currCourse.ID)
		fmt.Println("Name is : ", currCourse.Name)
	}
	
	uow.Commit()

}

func UpdateCourse() {

	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, false)
	defer uow.Complete()
	var currCourse model.Course
	currCourse.ID ,_= uuid.FromString("1d4512cd-e6d6-4503-afc2-e6c1747da4f7")
	currCourse.Name = "sanskrit"
	err := serviceInstanceCourse.gormRepo.Update(uow, &currCourse)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated successfully")
	}

	uow.Commit()

}


func DeleteCourse() {

	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, false)
	defer uow.Complete()
	var currCourse model.Course
	currCourse.ID, _= uuid.FromString("1d4512cd-e6d6-4503-afc2-e6c1747da4f7")

	err := serviceInstanceCourse.gormRepo.Delete(uow, &currCourse)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\nDeleted user successfully")
	}

	/*var hobby model.Hobby
	hobby.UserID=user.ID
	err := serviceInstanceCourse.gormRepo.Delete(uow, &hobby)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted hobbies successfully")
	}

	Isse pura user he gayab ho gaya hai
	*/
	uow.Commit()

}
