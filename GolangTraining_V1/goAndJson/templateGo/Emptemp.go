package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

type properties struct{
	Name string
	Dtype string
}

type newStruct struct{
	StructName string
	StructProperty []properties
}

func createMicroservice(w http.ResponseWriter,r *http.Request){
	props := []properties{
							{Name : "EmpName", Dtype : "string"},
							{Name : "EmailID", Dtype: "int"},
							{Name : "EmpSal",Dtype: "float32"},
						}

	emp := newStruct{
						StructName: "Employee",
						StructProperty: props,
					}
	
	templatePath := "C:\\Users\\aishwarya.anand\\OneDrive - Forcepoint\\Desktop\\Swabhav_Techlabs\\GolangTraining\\goAndJson\\templateGo\\structTemplate.txt"

	tmpl,err:= template.New("structTemplate.txt").ParseFiles(templatePath)
	if err!=nil{
		log.Fatal(err)
	}

	filename := emp.StructName + ".go"
	f,err := os.Create(filename)

	if err!=nil{
		log.Println("Create file error : ",err)
		return
	}
	tmpl.Execute(f,emp)
}

func initrouter(){
	r := mux.NewRouter()
	r.HandleFunc("/createMS",createMicroservice).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090",r))
}

func main(){
	initrouter()
}