package service

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	
	"gorm/model"
	"gorm/repository"
	"gorm/web"
)


func AddCourse(w http.ResponseWriter, r *http.Request){

	var serviceInstanceCourse = getInstanceOfService()

	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, false)
	defer uow.Complete()

	var singleCourse model.Course
	err := web.UnmarshalJSON(r,&singleCourse)
	if err!=nil{
		web.RespondErrorMessage(w,http.StatusBadRequest,err.Error())
		return
	}
	
	err = serviceInstanceCourse.gormRepo.Add(uow, &singleCourse)
	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	}

	web.RespondJSON(w,http.StatusCreated,"Course added successfully")
	uow.Commit()
}


func GetCourse(w http.ResponseWriter, r *http.Request) {
	var serviceInstanceCourse = getInstanceOfService()

	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, true)
	defer uow.Complete()

	var singleCourse model.Course
	var preloadAssoc = []string{}
	vars := mux.Vars(r)
	singleCourse.ID ,_= uuid.FromString(vars["id"])

	qp := repository.Filter("id = ?",singleCourse.ID)
	
	err := serviceInstanceCourse.gormRepo.Get(uow, &singleCourse, singleCourse.ID, preloadAssoc,qp)
	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	}

	web.RespondJSON(w,http.StatusOK,singleCourse)
	uow.Commit()

}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	var serviceInstanceCourse = getInstanceOfService()

	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, true)
	defer uow.Complete()

	var courses []model.Course
	err := serviceInstanceCourse.gormRepo.GetAll(uow, &courses, []string{})
	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	}
	
	web.RespondJSON(w,http.StatusOK,courses)
	uow.Commit()

}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var serviceInstanceCourse = getInstanceOfService()


	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, true)
	defer uow.Complete()

	vars := mux.Vars(r)
	var singleCourse model.Course
	singleCourse.ID ,_= uuid.FromString(vars["id"])
	err := web.UnmarshalJSON(r,&singleCourse)
	if err!=nil{
		web.RespondErrorMessage(w,http.StatusBadRequest,err.Error())
		return
	}

	err = serviceInstanceCourse.gormRepo.Update(uow, &singleCourse)
	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	} 
		
	web.RespondJSON(w,http.StatusOK,"Updated course successfully")
	uow.Commit()

}


func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	var serviceInstanceCourse = getInstanceOfService()
	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, false)
	defer uow.Complete()
	
	vars := mux.Vars(r)
	var currCourse model.Course
	currCourse.ID,_ = uuid.FromString(vars["id"])
	err := serviceInstanceCourse.gormRepo.Delete(uow, &currCourse)
	if err != nil {
		web.RespondErrorMessage(w,http.StatusInternalServerError,err.Error())
		return
	} 

	web.RespondJSON(w,http.StatusOK,"Course user successfully")
	uow.Commit()


}
