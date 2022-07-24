

package main

import (
	
	"bufio"
	"os"
	"strings"
	"wallet/service"
	"fmt"
)


func main() {
	walletService := service.GetInstance()
	/* 
	A Facade class can often be transformed into a Singleton since a single facade object is sufficient in most cases.
	Example,just one WalletService facade
	*/

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
		walletServicee := service.GetInstance()  
		/* Single instance already created.
		 the same singleton instance is returned every time by the GetInstance()
		*/
		accName,codeByUser := takeInput()
		err := walletServicee.Withdraw(accName,codeByUser)
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
