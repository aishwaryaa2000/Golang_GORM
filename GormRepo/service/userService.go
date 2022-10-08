package service

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	
	"gorm/encrypt"
	"gorm/authentication"
	"gorm/model"
	"gorm/repository"
	"gorm/web"
)

type UserLogin struct {
	Id string `json:"id"`
	Password   string `json:"password"`
}


func Login(w http.ResponseWriter, r *http.Request){
	var serviceInstanceUser = getInstanceOfService()

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, true)
	defer uow.Complete()

	var currentLogin UserLogin	
	err := web.UnmarshalJSON(r,&currentLogin)
	if err!=nil{
		web.RespondErrorMessage(w,http.StatusBadRequest,err.Error())
		return
	}

	var userr model.User
	var preloadAssoc = []string{"Courses", "Hobbies"}
	currID ,_ := uuid.FromString(currentLogin.Id)
	hashedPassword := encrypt.CreateHash(currentLogin.Id + currentLogin.Password)

	qp := repository.Filter("id = ? AND password = ? ",currID,hashedPassword)
	err = serviceInstanceUser.gormRepo.Get(uow, &userr, currID, preloadAssoc,qp)

	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return

	}

	token,_ := authentication.GenerateJWT(currentLogin.Id)
	authentication.SetCookieValue(w,token)
	web.RespondJSON(w,http.StatusOK,"Successfully logged in.")
	uow.Commit()

}

func AddUser(w http.ResponseWriter, r *http.Request) {

	var serviceInstanceUser = getInstanceOfService()

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()

	var user model.User
	err := web.UnmarshalJSON(r,&user)
	if err!=nil{
		web.RespondErrorMessage(w,http.StatusBadRequest,err.Error())
		return
	}

	var omit = []string{"Courses.*"}
	//New record in course table will not be added here,it has to be done via courseService.go
	err = serviceInstanceUser.gormRepo.AddWithOmit(uow, &user,omit)
	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	}
	web.RespondJSON(w,http.StatusCreated,"User added successfully")
	uow.Commit()
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	var serviceInstanceUser = getInstanceOfService()

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
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	}

	web.RespondJSON(w,http.StatusOK,user)
	uow.Commit()

}

func GetAllUser(w http.ResponseWriter, r *http.Request) {

	var serviceInstanceUser = getInstanceOfService()

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, true)
	defer uow.Complete()

	var users []model.User
	var preloadAssoc = []string{"Courses", "Hobbies"}
	err := serviceInstanceUser.gormRepo.GetAll(uow, &users, preloadAssoc)
	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	}
	
	web.RespondJSON(w,http.StatusOK,users)
	uow.Commit()

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var serviceInstanceUser = getInstanceOfService()

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()

	var user model.User
	user.ID = authentication.GetIdFromCookieClaims(r)

	err := web.UnmarshalJSON(r,&user)
	if err!=nil{
		web.RespondErrorMessage(w,http.StatusBadRequest,err.Error())
		return
	}

	if user.Password!=""{
		user.Password = encrypt.CreateHash(user.ID.String() + user.Password)
	}

	var omit = []string{"Courses.*"}
	//Dont add anything in courses table
	err = serviceInstanceUser.gormRepo.UpdateWithOmit(uow, &user,omit)
	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	} 

	web.RespondJSON(w,http.StatusOK,"Updated user successfully")
	uow.Commit()

}


func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var serviceInstanceUser = getInstanceOfService()

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()

	var user model.User
	user.ID = authentication.GetIdFromCookieClaims(r)

	err := serviceInstanceUser.gormRepo.Delete(uow, &user)
	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	} 
	authentication.DeleteCookieValue(w)
	web.RespondJSON(w,http.StatusOK,"Deleted user successfully")

	uow.Commit()

}

func DeleteHobbyOfAUser(w http.ResponseWriter, r *http.Request) {

	var serviceInstanceUser = getInstanceOfService()

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()
	

	userId := authentication.GetIdFromCookieClaims(r)

	var singleHobby model.Hobby

	err := web.UnmarshalJSON(r,&singleHobby)
	if err!=nil{
		web.RespondErrorMessage(w,http.StatusBadRequest,err.Error())
		return
	}

	singleHobby.UserID = userId
	qp := repository.Filter("name = ? AND user_id = ? ",singleHobby.Name,userId)
	err = serviceInstanceUser.gormRepo.Get(uow, &singleHobby, userId, []string{} ,qp)
	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	} 

	err = serviceInstanceUser.gormRepo.Delete(uow, &singleHobby)
	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	}

	web.RespondJSON(w,http.StatusOK,"Deleted hobby successfully")
	uow.Commit()

}


func DeleteCourseOfAUser(w http.ResponseWriter, r *http.Request){

	var serviceInstanceUser = getInstanceOfService()

	uow := repository.NewUnitOfWork(serviceInstanceUser.db, false)
	defer uow.Complete()
	

	vars := mux.Vars(r)
	var course model.Course
	course.ID,_ = uuid.FromString(vars["courseid"])

	var user model.User
	user.ID = authentication.GetIdFromCookieClaims(r)

	err := serviceInstanceUser.gormRepo.RemoveAssociations(uow,&user,"Courses",&course)
	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	}

	web.RespondJSON(w,http.StatusOK,"Deleted course mapped to the user successfully")
	uow.Commit()

}

