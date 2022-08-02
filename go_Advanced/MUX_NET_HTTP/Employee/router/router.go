package router

import(
	"Employee/api"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func MuxRoute(){
	r := mux.NewRouter()
	r.HandleFunc("/EmployeeRegister", api.EmployeeRegister).Methods("POST")
	r.HandleFunc("/UpdateEmployeeById", api.UpdateEmployeeById).Methods("POST")
	r.HandleFunc("/DeleteEmployeeById", api.DeleteEmployeeById).Methods("POST")
	r.HandleFunc("/DisplayEmployeeById", api.DisplayEmployeeById).Methods("POST")
	
	fmt.Println("Server started")
	http.ListenAndServe(":9090", r)
}
// gorm for extracting variables via url
// https://content-www.enterprisedb.com/postgres-tutorials/postgresql-and-golang-tutorial
  