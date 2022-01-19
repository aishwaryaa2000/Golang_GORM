package main

import (
	"accountApp/account"
	"accountApp/user"
	"accountApp/files"
	"fmt"
)

func main() {
	files.ReadFromFile()
begin:
	fmt.Printf("1-List all users\n2-Login\n3-Register a new user\n4-Exit\nEnter your choice : ")
	var userChoice int
	fmt.Scanln(&userChoice)
	switch userChoice {
	case 1:
		err:= user.DisplayAllUser()
		if err!=nil{
			fmt.Println(err)
		}
		goto begin //if I don't write goto then main() will exit
	case 2:
		var loginId int
		var pass string
		fmt.Print("\nEnter login ID : ")
		fmt.Scanln(&loginId)
		fmt.Print("Enter password : ")
		fmt.Scanln(&pass)
		currUser, err := user.Login(loginId, pass)
		if err != nil {
			fmt.Println("Error occurred : ",err)
			goto begin
		}
		fmt.Println("Login successfull")
		afterLoginMenu(currUser)
		goto begin //if I don't write goto then after LOGOUT program will exit
	case 3:
		var name, pass string
		fmt.Print("\nEnter your name : ")
		fmt.Scanln(&name)
		fmt.Print("Enter your password : ")
		fmt.Scanln(&pass)
		var newUser = user.New(name, pass,0)
		fmt.Println("Succesfully added")
		fmt.Println("Your user ID is : ", newUser.Id)
		files.WriteNewUser(newUser.Id,name,pass)
		goto begin
	case 4:
		return

	}

}

func afterLoginMenu(userLogin *user.User) {
LoginDashboard:
	fmt.Printf("1-Add account\n2-Transfer to an account\n3-Deposit to an account")
	fmt.Printf("\n4-Withdraw from an account\n5-Logout\nEnter your choice : ")
	var userChoice int
	fmt.Scanln(&userChoice)
	switch userChoice {
	case 1:
		var newAcc = account.New(0,0) //retuurns pointer- 0,0 for file handling part
		userLogin.AddAccount(newAcc)
		goto LoginDashboard //if I don't write goto then execution will go to the func main()
	case 2:
		//choose from which account do u wish to transfer
		var accNumber int
		found := false
		if userLogin.AllAccounts==nil{
			fmt.Println("You do not have any accounts to make a transfer")
			goto LoginDashboard
		}
		fmt.Print("Enter your account number : ")
		fmt.Scanln(&accNumber)
		for _, iaccount := range userLogin.AllAccounts {
			if iaccount.AccNo == accNumber {
				found = true
				var toAccNo, amt int
				fmt.Printf("To which account you wish to transfer? ")
				fmt.Scanln(&toAccNo)
				fmt.Printf("\nAmount you wish to transfer? ")
				fmt.Scanln(&amt)
				str, err := iaccount.Transfer(toAccNo, amt)
				if err != nil {
					fmt.Println("Error occurred : ",err)
					goto LoginDashboard				
				}
				fmt.Println(str)
				break
			}
		}
		if found == false {
			fmt.Println("Account number is incorrect")

		}
		goto LoginDashboard				


	case 3:
		//choose to which account u wish to deposit bcoz a user has multiple accounts
		var accNumber int
		if userLogin.AllAccounts == nil {
			fmt.Println("You don't have an account.Create one")
			goto LoginDashboard
		}
		found := false
		fmt.Print("Enter your account number : ")
		fmt.Scanln(&accNumber)
		for _, iaccount := range userLogin.AllAccounts {
			if iaccount.AccNo == accNumber {
				found = true
				var amt int
				fmt.Printf("\nAmount you wish to deposit? ")
				fmt.Scanln(&amt)
				str, err := iaccount.Deposit(amt)
				if err != nil {
					fmt.Println("Error occurred : ",err)
					goto LoginDashboard				
				}
				fmt.Println(str)
				break
			}
		}
		if found == false {
			fmt.Println("Your account number is incorrect")

		}
		goto LoginDashboard

	case 4:
		//choose from which account you wish to withdraw
		var accNumber int
		if userLogin.AllAccounts==nil{
			fmt.Println("You do not have any accounts to make a transfer")
			goto LoginDashboard
		}
		found := false
		fmt.Print("Enter your account number : ")
		fmt.Scanln(&accNumber)
		for _, iaccount := range userLogin.AllAccounts {
			if iaccount.AccNo == accNumber {
				found = true
				var amt int
				fmt.Printf("\nAmount you wish to withdraw? ")
				fmt.Scanln(&amt)
				str, err := iaccount.Withdraw(amt)
				if err != nil {
					fmt.Println("Error occurred : ",err)
					goto LoginDashboard
				}
				fmt.Println(str)
				break
			}
		}
		if found == false {
			fmt.Println("Your account number is incorrect")
		}
		goto LoginDashboard

	case 5:
		return

	}

}
