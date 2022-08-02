
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Users struct{
	U []User `json:"Users"`
}

type User struct{
	UserName string `json:"Name"`
	UserPhNumber Phone `json:"phoneNumbers"`
	UserAge uint8 `json:"Age"`
	UserPets []Pet `json:"Pets"` //slice of struct pets
	UserAddress Address `json:"Address"`
}

type Phone struct{
	HomeNo string `json:"home"`
	OfficeNo string `json:"office"`
}

type Pet struct{
	Type string `json:"Type"`
	Name string `json:"Name"`
}

type Address struct{
	Permanent string `json:"Permanent address"`
	Current string `json:"current Address"`
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/createEMP",createEMP).Methods("POST") 
	// POST is used to send data to a server to create/update a resource. 
	// The data sent to the server with POST is stored in the request body of the HTTP request
	fmt.Println("Sever started")
	http.ListenAndServe(":9090",r)
}
func createEMP(w http.ResponseWriter,r *http.Request){
	//r is request from client to server
	var mappedData Users
	jsonByteData,err2 := ioutil.ReadAll(r.Body) 
	//r.Body is the request body given from client to server
	if err2!=nil{
		fmt.Println(err2)
	}
	json.Unmarshal(jsonByteData,&mappedData)
	//Displaying the data given by the client 
	//Here,the data is given by postman POST method under which body consists of JSON data
	for _,iUser := range mappedData.U{
		fmt.Println("Name : ",iUser.UserName)
		fmt.Println("Phone numbers \n    Home number",iUser.UserPhNumber.HomeNo,"\n    Office number : ",iUser.UserPhNumber.OfficeNo)
		fmt.Println("Age : ",iUser.UserAge)
		fmt.Println("Pets")
		for _,iPet := range iUser.UserPets{
			fmt.Println("     Type : ",iPet.Type,"\n     Name : ",iPet.Name)
		}
		fmt.Println("Address \n    Permanent address : ",iUser.UserAddress.Permanent,"\n    Current Address : ",iUser.UserAddress.Current)
	}
}