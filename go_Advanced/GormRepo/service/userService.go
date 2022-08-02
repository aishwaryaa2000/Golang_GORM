package service

import (
	"fmt"
	"gorm/model"
	"gorm/repository"
	"github.com/satori/go.uuid"
)


var serviceInstanceUser = getInstanceOfService()


func AddUser() {

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()
	id:=uuid.NewV4()
	// course1 := model.NewCourse("History")
	// course2 := model.NewCourse("Politcs")
	hobby1 := model.NewHobby("skating")
	hobby2 := model.NewHobby("acting")
	courseID ,_ := uuid.FromString("712ca4b1-8c35-4871-9673-45bd49a6394a")
	//here with the help of ID,I can fetch the course name from course table
	courseName := "Geography"
	userr := model.User{
		FName: "Ramani",
		LName: "Anand",
		CustomModel: model.CustomModel{
			ID : id,
		},
		Courses: []model.Course{
			{
				CustomModel : model.CustomModel{
					ID :courseID,
				},
				Name: courseName,
			},
		
		},
		Hobbies: []model.Hobby{
			hobby1,
			hobby2,
		},
	}

	/*
    For many2many associations, GORM will upsert the associations before creating the join table references,
	if you want to skip the upserting of associations, you could skip it like:
    db.Omit("Courses.*").Create(&user)
	Values will get added in user_courses but not in courses table
	courses table me buss name update hoga according to courseID specified while creation

    The following code will skip the creation of the association and its references
	db.Omit("Courses").Create(&user)
	In this case,values won't get added in user_course and courses table
	*/

	var omit = []string{"Courses.*"}
	//New record in course table will not be added here,it has to be done via courseService.go
	err := serviceInstanceUser.gormRepo.AddWithOmit(uow, &userr,omit)
	if err != nil {
		fmt.Println(err)
	}

	uow.Commit()
}

func GetUser() {

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, true)
	defer uow.Complete()

	var userr model.User
	var preloadAssoc = []string{"Courses", "Hobbies"}
	currID ,_ := uuid.FromString("a5903424-cd3e-4025-abdb-e40084ddc7b9")
	err := serviceInstanceUser.gormRepo.Get(uow, &userr, currID, preloadAssoc)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ID is : ", userr.ID)
	fmt.Println("Name is : ", userr.FName, userr.LName)
	fmt.Print("Hobbies are ")
	for _, iHobby := range userr.Hobbies {
		fmt.Print(iHobby.Name, " ")
	}
	fmt.Print("\nCourses enrolled are ")
	for _, iCourse := range userr.Courses {
		fmt.Print(iCourse.Name, " ")
	}

	uow.Commit()

}

func GetAllUser() {

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, true)
	defer uow.Complete()

	var users []model.User
	var preloadAssoc = []string{"Courses", "Hobbies"}
	err := serviceInstanceUser.gormRepo.GetAll(uow, &users, preloadAssoc)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nLIST OF ALL USERS")
	for _, currUser := range users {
		fmt.Println("\nID is : ", currUser.ID)
		fmt.Println("Name is : ", currUser.FName, currUser.LName)
		fmt.Print("Hobbies : ")
		for _, iHobby := range currUser.Hobbies {
			fmt.Print(iHobby.Name, " ")
		}
		fmt.Print("\nCourses enrolled : ")
		for _, iCourse := range currUser.Courses {
			fmt.Print(iCourse.Name, " ")
		}
	}

	uow.Commit()

}

func UpdateUser() {

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()
	var user model.User
	user.ID ,_= uuid.FromString("dbbe9842-be7b-451f-bd1c-2206b9bfbbd6")
	user.FName = "Aishu"
	user.LName = "Kutty"
	err := serviceInstanceUser.gormRepo.Update(uow, &user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated successfully")
	}

	uow.Commit()

}


func DeleteUser() {

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()
	var user model.User
	user.ID, _= uuid.FromString("dbbe9842-be7b-451f-bd1c-2206b9bfbbd6")

	err := serviceInstanceUser.gormRepo.Delete(uow, &user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\nDeleted user successfully")
	}

	/*var hobby model.Hobby
	hobby.UserID=user.ID
	err := serviceInstanceUser.gormRepo.Delete(uow, &hobby)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted hobbies successfully")
	}

	Isse pura user he gayab ho gaya hai
	*/
	uow.Commit()

}



func HardDeleteUser() {

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()
	var user model.User
	user.ID, _= uuid.FromString("0d1ddf9f-53e2-4440-9545-d72679582fed")

	err := serviceInstanceUser.gormRepo.HardDelete(uow, &user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\nDeleted user successfully")
	}

	/*var hobby model.Hobby
	hobby.UserID=user.ID
	err := serviceInstanceUser.gormRepo.Delete(uow, &hobby)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted hobbies successfully")
	}

	Isse pura user he gayab ho gaya hai
	*/
	uow.Commit()

}