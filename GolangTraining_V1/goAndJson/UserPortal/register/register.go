package register

import (
	"UserPortal/connect"
	"UserPortal/loginSchema"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"
)
func isValidEmail(email string)bool{
	_,err := mail.ParseAddress(email)
	return err == nil
}

func IsValidate(Email,Password,ConfirmPassword string) (error){
	if isValidEmail(Email)==false{
		err := errors.New("Email ID is not in correct format")
		return err
	}else if(!(Password==ConfirmPassword)){
		err := errors.New("New password and confirm password is not same")
		return err
	}
	return nil	
}

func UserRegister(w http.ResponseWriter,r *http.Request){
	var mappedData loginSchema.LoginSchema
	jsonByteData,err2 := ioutil.ReadAll(r.Body) 
	//r.Body is the request body given from client to server
	if err2!=nil{
		fmt.Println(err2)
	}
	json.Unmarshal(jsonByteData,&mappedData)
	
	err := IsValidate(mappedData.Email,mappedData.Password,mappedData.ConfirmPassword)
	if(err!=nil){
		fmt.Println("Error occured : ",err)
	}else{
		fmt.Println("Username for registration : ",mappedData.Username,"and password : ",mappedData.Password)
		queryString := `INSERT INTO logindb (username,passwrd,email,dob) VALUES('` + mappedData.Username + `','` + mappedData.Password +  `','` + mappedData.Email + `','` + mappedData.Dob + `')`

		db := connect.SetupDB()
		_,err := db.Exec(queryString)
		if err!=nil{
			fmt.Println(err)
		}else{
			fmt.Println("Registered successfully")
		}
	}

}