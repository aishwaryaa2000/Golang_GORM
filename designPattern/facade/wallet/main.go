package main

import (
	"wallet/components/account"
	"wallet/components/ledger"
	"wallet/components/securitycode"
	"wallet/components/wallet"
	"bufio"
	"os"
	"strings"
	"wallet/components/notification"
	"errors"
	"fmt"
)


var l ledger.Ledger

func main() {
	accountCurrent := account.New()
	walletCurrent := wallet.New()

begin:
	fmt.Printf("\nMENU\n1)-Deposit\n2)-Withdraw\n3)-Get all ledger entries\n4)-Exit\nEnter your choice : ")
	var userChoice int
	fmt.Scanln(&userChoice)
	switch userChoice {
	case 1:
		err := Deposit(accountCurrent,walletCurrent)
		if err != nil {
			fmt.Println(err)
		}
		goto begin
	case 2:
		err := Withdraw(accountCurrent,walletCurrent)
		if err != nil {
			fmt.Println(err)
		}
		goto begin

	case 3:
		l.GetAllEntries()
		goto begin
	case 4:
		return
	default:
		fmt.Println("Incorrect choice")
		goto begin

	}
}

func Deposit(accountCurrent *account.Account,walletCurrent * wallet.Wallet) error {
	in := bufio.NewReader(os.Stdin)

	securityCodeCurrent := securitycode.New()
	var accName string
	var codeByUser string
	fmt.Println("Enter account name")
	accName, _ = in.ReadString('\n')
	accName = strings.TrimRight(accName, "\r\n")
	fmt.Println("Enter check code")
	codeByUser, _ = in.ReadString('\n')
	codeByUser = strings.TrimRight(codeByUser, "\r\n")

	if accountCurrent.CheckAccount(accName) && securityCodeCurrent.CheckCode(codeByUser) { //true
		amtDeposit := 0
		fmt.Println("Enter amount to be deposited")
		fmt.Scanln(&amtDeposit)
		walletCurrent.Credit(amtDeposit)
		l.MakeEntry(accName, "deposited", amtDeposit)
		n := notification.Notification{}
		notificationMsg := n.SendCreditNotification()
		fmt.Println(notificationMsg)
		return nil
	}

	err := errors.New("incorrect name or code")
	return err

}

func Withdraw(accountCurrent *account.Account,walletCurrent * wallet.Wallet) error {
	in := bufio.NewReader(os.Stdin)

	securityCodeCurrent := securitycode.New()
	var accName string
	var codeByUser string
	fmt.Println("Enter account name")
	accName, _ = in.ReadString('\n')
	accName = strings.TrimRight(accName, "\r\n")
	fmt.Println("Enter check code")
	codeByUser, _ = in.ReadString('\n')
	codeByUser = strings.TrimRight(codeByUser, "\r\n")

	if accountCurrent.CheckAccount(accName) && securityCodeCurrent.CheckCode(codeByUser) { //true
		amtWithdraw := 0
		withdraw:
		fmt.Println("Enter amount to be writhdrawn")
		fmt.Scanln(&amtWithdraw)
		err := walletCurrent.Debit(amtWithdraw)
		if err != nil {
			fmt.Println(err)
			goto withdraw
		}
		l.MakeEntry(accName, "withdrawal", amtWithdraw)
		n := notification.Notification{}
		notificationMsg := n.SendDebitNotification()
		fmt.Println(notificationMsg)
		return nil
	}

	err := errors.New("incorrect name or code")
	return err

}
