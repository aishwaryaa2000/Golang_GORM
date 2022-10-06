package service

import (
	"encoding/json"
	"fmt"
	"gorm/encrypt"
	"gorm/jwtauthentication"
	"gorm/model"
	"gorm/repository"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

type UserLogin struct {
	Id string `json:"id"`
	Password   string `json:"password"`
}

var serviceInstanceUser = getInstanceOfService()

func Login(w http.ResponseWriter, r *http.Request){

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, true)
	defer uow.Complete()

	var currentLogin UserLogin	
	err := UnmarshalJSON(r,&currentLogin)
	if err!=nil{
		w.Write([]byte(err.Error()))
		return
	}

	var userr model.User
	var preloadAssoc = []string{"Courses", "Hobbies"}
	currID ,_ := uuid.FromString(currentLogin.Id)
	hashedPassword := encrypt.CreateHash(currentLogin.Id + currentLogin.Password)
	qp := repository.Filter("id = ? AND password = ? ",currID,hashedPassword)
	
	err = serviceInstanceUser.gormRepo.Get(uow, &userr, currID, preloadAssoc,qp)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Username or Password is incorrect")
		return

	}

	token,_ := jwtauthentication.GenerateJWT(currentLogin.Id)
	outputString :="Successfully logged in.Token is : " +token

	http.SetCookie(w,&http.Cookie{
		Name : "token",
		Value : token,
		Expires: time.Now().Add(time.Minute*5),
	})

	json.NewEncoder(w).Encode(outputString)

	uow.Commit()

}

func AddUser(w http.ResponseWriter, r *http.Request) {

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()

	var user model.User
	err := UnmarshalJSON(r,&user)
	if err!=nil{
		w.Write([]byte(err.Error()))
		return
	}

	var omit = []string{"Courses.*"}
	//New record in course table will not be added here,it has to be done via courseService.go
	err = serviceInstanceUser.gormRepo.AddWithOmit(uow, &user,omit)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error occured while registering the user"))
		return
	}
	w.Write([]byte("User added successfully"))
	uow.Commit()
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, true)
	defer uow.Complete()

	vars := mux.Vars(r)
	id := vars["id"]

	var user model.User
	user.ID ,_= uuid.FromString(id)
	var preloadAssoc = []string{"Courses", "Hobbies"}
	
	qp := repository.Filter("id = ?",user.ID)
	
	err := serviceInstanceUser.gormRepo.Get(uow, &user, user.ID, preloadAssoc,qp)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error while getting the user with specifies user id"))
		return
	}

	outputString := "\nID is : " + user.ID.String() + "\nName is : " + user.FName + " " + user.LName + "\n Hobbies are : "
	
	for _, iHobby := range user.Hobbies {
		outputString += iHobby.Name + " "
	}
	outputString +=  "\nCourses enrolled are "
	for _, iCourse := range user.Courses {
		outputString += iCourse.Name + " "
	}
	w.Write([]byte(outputString))
	uow.Commit()

}

func GetAllUser(w http.ResponseWriter, r *http.Request) {

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, true)
	defer uow.Complete()

	var users []model.User
	var preloadAssoc = []string{"Courses", "Hobbies"}
	err := serviceInstanceUser.gormRepo.GetAll(uow, &users, preloadAssoc)
	if err != nil {
		fmt.Println(err)
	}
	
	outputString:= "LIST OF ALL USERS\n"
	for _, currUser := range users {
		outputString += "\n\nID is : " +  currUser.ID.String() + "\nName is : " + currUser.FName + " " +currUser.LName + "\nHobbies : "
		
		for _, iHobby := range currUser.Hobbies {
			outputString += iHobby.Name + " "
		}
		outputString += "\nCourses : "
		for _, iCourse := range currUser.Courses {
			outputString += iCourse.Name + " "
		}
	}

	w.Write([]byte(outputString))

	uow.Commit()

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {


	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()

	var user model.User
	user.ID = jwtauthentication.GetIdFromCookieClaims(r)
	err := UnmarshalJSON(r,&user)
	if err!=nil{
		w.Write([]byte(err.Error()))
		return
	}

	if user.Password!=""{
		user.Password = encrypt.CreateHash(user.ID.String() + user.Password)
	}

	var omit = []string{"Courses.*"}
	//Dont add anything in courses table
	err = serviceInstanceUser.gormRepo.UpdateWithOmit(uow, &user,omit)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error occured while updating the user"))
		return
	} 

	w.Write([]byte("\nUpdated user successfully"))

	uow.Commit()

}


func DeleteUser(w http.ResponseWriter, r *http.Request) {

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()

	var user model.User
	user.ID = jwtauthentication.GetIdFromCookieClaims(r)

	err := serviceInstanceUser.gormRepo.Delete(uow, &user)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error occured while deleting the user"))
		return
	} 

	http.SetCookie(w,&http.Cookie{
		Name : "token",
		MaxAge:   -1,
		// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	})

	w.Write([]byte("\nDeleted user successfully"))
		
	uow.Commit()

}

func DeleteHobbyOfAUser(w http.ResponseWriter, r *http.Request) {


	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()
	

	userId := jwtauthentication.GetIdFromCookieClaims(r)

	var singleHobby model.Hobby

	err := UnmarshalJSON(r,&singleHobby)
	if err!=nil{
		w.Write([]byte(err.Error()))
		return
	}

	singleHobby.UserID = userId
	qp := repository.Filter("name = ? AND user_id = ? ",singleHobby.Name,userId)
	err = serviceInstanceUser.gormRepo.Get(uow, &singleHobby, userId, []string{} ,qp)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("No such hobby found"))
		return
	} 

	err = serviceInstanceUser.gormRepo.Delete(uow, &singleHobby)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error while deleting the hobby"))
		return
	}

	w.Write([]byte("Deleted hobby successfully"))
	

	uow.Commit()

}


func DeleteCourseOfAUser(w http.ResponseWriter, r *http.Request){

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()
	

	vars := mux.Vars(r)
	var course model.Course
	course.ID,_ = uuid.FromString(vars["courseid"])

	var user model.User
	user.ID = jwtauthentication.GetIdFromCookieClaims(r)

	err := serviceInstanceUser.gormRepo.RemoveAssociations(uow,&user,"Courses",course)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error while deleting the course"))
		return
	}

	w.Write([]byte("Deleted course mapped to the user successfully"))
	uow.Commit()

}

