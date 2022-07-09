package user

import (
	"accountApp/account"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var allUserSlice []*User

type User struct {
	Id          int
	name        string
	password    string
	AllAccounts []*account.Account
}

func(currUser *User) ReturnNameAndPass() (string,string){
	return currUser.name, currUser.password
}

func New(name, password string,id int) *User {
	var userNew = User{
		name:     name, //fieldName : VariableName
		password: password,
		Id:       getNewId(id),
	}
	allUserSlice = append(allUserSlice, &userNew)
	return &userNew //this is a pointer
}

func getNewId(id int) int {

	if id!=0{
		return id //if id is passed via file
	}	
begin:
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	newId := rand.Intn(100) + 1
	getUser := GetAllUsers()
	for _, value := range getUser {
		if value.Id == newId {
			goto begin
		}
	}

	return newId

}

func GetAllUsers() []*User {

	return allUserSlice

}

func DisplayAllUser() error{
	if allUserSlice==nil{
		return errors.New("No users to display")
	}
	for _, iuser := range allUserSlice {
		fmt.Println("\nUser ID is : ", iuser.Id)
		fmt.Println("User name is : ", iuser.name)
	
		for _, iaccount := range iuser.AllAccounts {
			fmt.Println("Account ID is : ", iaccount.AccNo)
			fmt.Println("Account Balance is : ", iaccount.Balance)

		}
	}
	return nil
}

func Login(id int, pass string) (*User, error) {
	for _, iuser := range allUserSlice {
		if iuser.Id == id && iuser.password == pass {
			return iuser, nil
		}
	}
	return nil, errors.New("Login failed.User ID with given password not found")

}

func (currUser *User) AddAccount(acc *account.Account) {

	currUser.AllAccounts = append(currUser.AllAccounts, acc)
	

}
