package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)



type joke struct{
	Categories []string `json:"categories"` 
	Created_at string `json:"created_at"`
	Icon_url string `json:"icon_url"`
	Id string `json:"id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Value string `json:"value"`

}


func main(){
	url := "https://api.chucknorris.io/jokes/random"
	res,err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println("\n\njson response body : ",string(body))


	joke1 := joke{}
	jsonErr := json.Unmarshal(body, &joke1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println("\n\nafter decoding using unmarshal : ",joke1)


	joke2 := joke{}

	errr := json.NewDecoder(res.Body).Decode(&joke2)
	if errr!=nil{
		fmt.Println(err)
	}

	fmt.Println("\n\nafter decoding using decoder : ",joke2)



}