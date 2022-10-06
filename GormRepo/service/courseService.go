package service

import (
	"fmt"
	"gorm/model"
	"gorm/repository"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)



func AddCourse(w http.ResponseWriter, r *http.Request){
	var serviceInstanceCourse = getInstanceOfService()

	
	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, false)
	defer uow.Complete()

	var singleCourse model.Course
	err := UnmarshalJSON(r,&singleCourse)
	if err!=nil{
		w.Write([]byte(err.Error()))
		return
	}
	
	err = serviceInstanceCourse.gormRepo.Add(uow, &singleCourse)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error occurred while adding a course"))
		return
	}
	w.Write([]byte("Course added successfully"))
	uow.Commit()
}


func GetCourse(w http.ResponseWriter, r *http.Request) {
	var serviceInstanceCourse = getInstanceOfService()


	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, true)
	defer uow.Complete()

	vars := mux.Vars(r)
	id := vars["id"]

	var singleCourse model.Course
	singleCourse.ID ,_= uuid.FromString(vars["id"])

	
	qp := repository.Filter("id = ?",singleCourse.ID)
	
	var currCourse model.Course
	var preloadAssoc = []string{}
	err := serviceInstanceCourse.gormRepo.Get(uow, &currCourse, singleCourse.ID, preloadAssoc,qp)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error occurred while getting the course"))
		return
	}

	outputString := "Course ID is : " + id +".Name of the course is : " + currCourse.Name
	w.Write([]byte(outputString))

	uow.Commit()

}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	var serviceInstanceCourse = getInstanceOfService()

	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, true)
	defer uow.Complete()

	var courses []model.Course
	err := serviceInstanceCourse.gormRepo.GetAll(uow, &courses, []string{})
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error occurred while getting courses"))
		return
	}

	outputString := "LIST OF ALL COURSES\n"
	for _, currCourse := range courses {
		outputString = outputString + "Course ID is : " + currCourse.ID.String() +".Name of the course is : " + currCourse.Name+"\n"
	}
	w.Write([]byte(outputString))

	uow.Commit()

}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var serviceInstanceCourse = getInstanceOfService()


	uow := repository.NewUnitOfWork(serviceInstanceCourse.db, true)
	defer uow.Complete()

	vars := mux.Vars(r)
	var singleCourse model.Course
	singleCourse.ID ,_= uuid.FromString(vars["id"])
	err := UnmarshalJSON(r,&singleCourse)
	if err!=nil{
		w.Write([]byte(err.Error()))
		return
	}

	err = serviceInstanceCourse.gormRepo.Update(uow, &singleCourse)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error occurred while updating course"))
		return
	} 
		
	w.Write([]byte("Updated successfully"))

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
		fmt.Println(err)
		w.Write([]byte("Error occurred while deleting course"))
		return
	} 

	w.Write([]byte("Course deleted successfully"))
	uow.Commit()


}
