package main

import (
	
	"bufio"
	"os"
	"strings"
	"wallet/service"
	"fmt"
)


func main() {
	walletService := service.NewWalletService()

begin:
	fmt.Printf("\nMENU\n1)-Deposit\n2)-Withdraw\n3)-Get all ledger entries\n4)-Exit\nEnter your choice : ")
	var userChoice int
	fmt.Scanln(&userChoice)
	switch userChoice {
	case 1:
		accName,codeByUser := takeInput()
		
		err := walletService.Deposit(accName,codeByUser)
		if err != nil {
			fmt.Println(err)
		}

		goto begin

	case 2:
		accName,codeByUser := takeInput()
		err := walletService.Withdraw(accName,codeByUser)
		if err != nil {
			fmt.Println(err)
		}

		goto begin

	case 3:
		walletService.GetAllLedgerEntries()
		goto begin

	case 4:
		return

	default:
		fmt.Println("Incorrect choice")
		goto begin

	}
}


func takeInput() (string,string){
	in := bufio.NewReader(os.Stdin)
	var accName string
	var codeByUser string
	fmt.Println("Enter account name")
	accName, _ = in.ReadString('\n')
	accName = strings.TrimRight(accName, "\r\n")
	fmt.Println("Enter check code")
	codeByUser, _ = in.ReadString('\n')
	codeByUser = strings.TrimRight(codeByUser, "\r\n")

	return accName,codeByUser
}
