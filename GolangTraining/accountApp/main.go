package main

import (
	"accountApp/account"
	"accountApp/user"
	"fmt"
	"log"
)

func main() {
begin:
	fmt.Printf("1-List all users\n2-Login\n3-Register a new user\n4-Exit\nEnter your choice : ")
	var userChoice int
	fmt.Scanln(&userChoice)
	switch userChoice {
	case 1:
		user.DisplayAllUser()
		goto begin
	case 2:
		var loginId int
		var pass string
		fmt.Print("Enter login ID : ")
		fmt.Scanln(&loginId)
		fmt.Print("\nEnter password : ")
		fmt.Scanln(&pass)
		currUser, err := user.Login(loginId, pass)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Login successfull")
		afterLoginMenu(currUser)
		goto begin
	case 3:
		var name, pass string
		fmt.Print("\nEnter your name : ")
		fmt.Scanln(&name)
		fmt.Print("\nEnter your password : ")
		fmt.Scanln(&pass)
		var newUser = user.New(name, pass)
		fmt.Println("Succesfully added")
		fmt.Println("Your user ID is : ", newUser.Id)
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
		var newAcc = account.New() //retuurns pointer
		userLogin.AddAccount(newAcc)
		goto LoginDashboard
	case 2:
		//choose from which account do u wish to transfer
		var accNumber int
	insidecase2:
		found := false
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
			goto insidecase2
		}

	case 3:
		//choose to which account u wish to deposit bcoz a user has multiple accounts
		var accNumber int
	insidecase3:
		found := false
		fmt.Print("Enter your account number : ")
		fmt.Scanln(&accNumber)

		if userLogin.AllAccounts == nil {
			fmt.Println("This user doesn't have an account.Create one")
			goto LoginDashboard
		}
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
			fmt.Println("Account number is incorrect")
			goto insidecase3
		}
		goto LoginDashboard

	case 4:
		//choose from which account you wish to withdraw
		var accNumber int
	insidecase4:
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
			fmt.Println("Account number is incorrect")
			goto insidecase4
		}
		goto LoginDashboard

	case 5:
		return

	}

}
