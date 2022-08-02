package api

import (
	"Employee/Database"
	"Employee/employeeModel"
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
	if  !(isValidEmail(Email)){
		err := errors.New("Email ID is not in correct format")
		return err
	}else if(!(Password==ConfirmPassword)){
		err := errors.New("New password and confirm password is not same")
		return err
	}
	return nil	
}

func EmployeeRegister(w http.ResponseWriter,r *http.Request){
	var mappedData employeeModel.EmployeeModel
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
		queryString := `INSERT INTO employee (empid,name,passwrd,email,dob) VALUES('` + mappedData.Id + `','` + mappedData.Name +  `','`  + mappedData.Password +  `','` + mappedData.Email + `','` + mappedData.Dob + `')`
		db := Database.SetupDB()
		_,err := db.Exec(queryString)
		if err!=nil{
			fmt.Println(err)
		}else{
			fmt.Println("Registered successfully")
		}
	}

}


func UpdateEmployeeById(w http.ResponseWriter,r *http.Request){
	var mappedData employeeModel.EmployeeModel
	jsonByteData,err2 := ioutil.ReadAll(r.Body) 
	//r.Body is the request body given from client to server
	if err2!=nil{
		fmt.Println(err2)
	}
	json.Unmarshal(jsonByteData,&mappedData)
	db := Database.SetupDB()
	if !(isValidEmail(mappedData.Email)){
		fmt.Println("Email ID is not in correct format")

	}else{
		sqlStatement := `UPDATE employee SET email = '` +mappedData.Email+`' WHERE empid ='`+mappedData.Id+`'`
		res, err := db.Exec(sqlStatement)
		if err != nil {
			fmt.Println(err)
		}
		noOfRowsAffected, errUpdate := res.RowsAffected()
		if errUpdate != nil {
			fmt.Println(errUpdate)
		}else{
			fmt.Println("Updated succesfully.Rows affected : ",noOfRowsAffected)
		}
	}
}


func DeleteEmployeeById(w http.ResponseWriter,r *http.Request){
	var mappedData employeeModel.EmployeeModel
	jsonByteData,err2 := ioutil.ReadAll(r.Body) 
	//r.Body is the request body given from client to server
	if err2!=nil{
		fmt.Println(err2)
	}
	json.Unmarshal(jsonByteData,&mappedData)
	db := Database.SetupDB()
	
	sqlStatement := `DELETE FROM employee WHERE empid ='` +mappedData.Id+`'`
	_, errDel := db.Exec(sqlStatement)
	if errDel != nil {
	fmt.Println(errDel)
	}else{
		fmt.Println("Deleted successfully")
	}

}

func DisplayEmployeeById(w http.ResponseWriter,r *http.Request){
	var mappedData employeeModel.EmployeeModel
	jsonByteData,err2 := ioutil.ReadAll(r.Body) 
	//r.Body is the request body given from client to server
	if err2!=nil{
		fmt.Println(err2)
	}

	json.Unmarshal(jsonByteData,&mappedData)
	db := Database.SetupDB()
	
	sqlStatement := `SELECT * FROM employee WHERE empid ='` +mappedData.Id+`'`
	rows, errDisplay := db.Query(sqlStatement)
	if errDisplay != nil {
	fmt.Println(errDisplay)
	}

	for rows.Next() {
        var id, name,passwrd, email, dob string
        if err := rows.Scan(&id, &name, &passwrd, &email, &dob); err != nil {
            fmt.Println(err)
        }
        fmt.Printf(" Id: %s\n Name: %s \n Password: %s \n Email: %s \n Date of birth: %s\n---\n\n", id, name, passwrd, email, dob)
    } 


}