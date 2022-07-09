package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	jsonFile,err1 := os.Open("userData.json")
	if err1!=nil{
		fmt.Println(err1)
	}
	var mappedData Users
	jsonByteData,err2 := ioutil.ReadAll(jsonFile)
	if err2!=nil{
		fmt.Println(err2)
	}
	json.Unmarshal(jsonByteData,&mappedData)

	for _,iUser := range mappedData.U{
		fmt.Println("\nName : ",iUser.UserName)
		fmt.Println("Phone numbers \n    Home number",iUser.UserPhNumber.HomeNo,"\n    Office number : ",iUser.UserPhNumber.OfficeNo)
		fmt.Println("Age : ",iUser.UserAge)
		fmt.Println("Pets")
		for _,iPet := range iUser.UserPets{
			fmt.Println("     Type : ",iPet.Type,"\n     Name : ",iPet.Name)
		}
		fmt.Println("Address \n    Permanent address : ",iUser.UserAddress.Permanent,"\n    Current Address : ",iUser.UserAddress.Current)
	}
}