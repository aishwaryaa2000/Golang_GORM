package user

import (
	"accountApp/account"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var allUserSclice []User

type User struct {
	Id          int
	name        string
	password    string
	AllAccounts []*account.Account
}

func New(name, password string) *User {
	var userNew = &User{
		name:     name, //fieldName : VariableName
		password: password,
		Id:       getNewId(),
	}
	allUserSclice = append(allUserSclice, *userNew)
	return userNew //this is a pointer
}

func getNewId() int {

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

func GetAllUsers() []User {

	return allUserSclice

}

func DisplayAllUser() {
	for _, iuser := range allUserSclice {
		fmt.Println("\nUser ID is : ", iuser.Id)
		fmt.Println("User name is : ", iuser.name)
		if iuser.AllAccounts != nil {
			fmt.Println("Accounts are there")
		}
		for _, iaccount := range iuser.AllAccounts {
			fmt.Println("Account ID is : ", iaccount.AccNo)
			fmt.Println("Account Balance is : ", iaccount.Balance)

		}
	}
}

func Login(id int, pass string) (*User, error) {
	for _, iuser := range allUserSclice {
		if iuser.Id == id && iuser.password == pass {
			return &iuser, nil
		}
	}
	return nil, errors.New("Login failed.User ID with given password not found")

}

func (currUser *User) AddAccount(acc *account.Account) {

	currUser.AllAccounts = append(currUser.AllAccounts, acc)
	if currUser.AllAccounts != nil {
		fmt.Println("New account with account ID ", acc.AccNo, "has been created")
	}

}
