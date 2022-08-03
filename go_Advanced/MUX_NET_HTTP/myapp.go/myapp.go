package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"encoding/json"
)
type MyStruct struct{
	Icon_url string `json:"icon_url"`
	Id string `json:"id"`
}

// func WelcomeMsgHandler(w http.ResponseWriter,r *http.Request){
// 	fmt.Println("Welcome aishwarya in console")
// 	w.Write([]byte(fmt.Sprintf("Welcomee on server")))
// }

func WelcomeMsgHandler2(w http.ResponseWriter,r *http.Request){
	fmt.Println("Welcome aishwarya in console")
	vars := mux.Vars(r)
	code := r.URL.Query().Get("code")
	id := vars["id"]
    w.WriteHeader(http.StatusAccepted)
	w.Header().Set("my key", "value")
    fmt.Fprintf(w, "ID is : %v\n and code is %v", id,code)
	// w.Write([]byte(fmt.Sprintf("Welcomee on server  ",id)))
}


func Yo(w http.ResponseWriter,r *http.Request){
	p := MyStruct{}
	p.Icon_url="this is icon url"
	p.Id="this is id"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
func main(){
	r := mux.NewRouter()
	// r.HandleFunc("/welcome",WelcomeMsgHandler).Methods(http.MethodGet)
	r.HandleFunc("/welcome/{id}",WelcomeMsgHandler2).Methods(http.MethodGet)

	r.HandleFunc("/yo",Yo).Methods(http.MethodGet)

	address := "127.0.0.1:9090" 
	server := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	  }
	fmt.Println("Sever started")
	log.Fatal(server.ListenAndServe())
	// http.ListenAndServe(":9090",r)
}


//response 
//then response as json
//call jokes api in 25times and note the time taken
//phele sync then go routines ke saath
//25 call ko kitna time laga,sync me and async me
//u can use sync.map