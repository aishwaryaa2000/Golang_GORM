package login

import (
	"UserPortal/connect"
	"UserPortal/loginSchema"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UserLogin(w http.ResponseWriter,r *http.Request){
	var mappedData loginSchema.LoginSchema
	loginData,_ := ioutil.ReadAll(r.Body)
	json.Unmarshal(loginData,&mappedData)
	queryString := `SELECT * FROM logindb where username='` + mappedData.Username + `' and passwrd='` + mappedData.Password + `'`

	db := connect.SetupDB()
	row,err := db.Query(queryString)
	//db.QueryRow
	if err!=nil{
		fmt.Println(err)
	}else{

		if row.Next(){
			fmt.Println("Login successfull")
		}else{
			fmt.Println("Login failed.Incorrect credentials")
		}
	}
}