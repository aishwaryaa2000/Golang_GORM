package service

import (
	"fmt"
	"wallet/components/account"
	"wallet/components/ledger"
	"wallet/components/securitycode"
	"wallet/components/wallet"
	"wallet/components/notification"
	"errors"
	"sync"

)

var lock = &sync.Mutex{}

var singleInstanceOfWalletService *WalletService


type WalletService struct {
    account      *account.Account
    wallet       *wallet.Wallet
    securityCode *securitycode.SecurityCode
    notification *notification.Notification
    ledger       *ledger.Ledger
}


func GetInstance() *WalletService {
	//singleInstanceOfWalletService (singleton) is the shared variable if there were multiple go routines trying to access them
	if singleInstanceOfWalletService == nil {
		//Check if singleInstanceOfWalletService is already created or not
        lock.Lock()
		/* 
		Mutex provides mutual exclusion to a resource, which means that only one goroutine can hold it at a time.
		Any goroutine must first acquire the mutex and then access the resource. 

		If a goroutine tries to acquire the mutex already held by another goroutine,
		it will be blocked and will wait for the mutex to be released.
		*/
        defer lock.Unlock()
        if singleInstanceOfWalletService == nil {
			/* 
			check for nil singleInstanceOfWalletService
			after the lock is acquired. This is to make sure that
			if more than one goroutine bypass the first check then only one goroutine
			is able to create the singleton instance otherwise each of the goroutine 
			will create its own instance of the single struct.
			*/
            fmt.Println("Creating single instance now.")
            singleInstanceOfWalletService = &WalletService{
				account:      account.New(),
				securityCode: securitycode.New(),
				wallet:       wallet.New(),
				notification: &notification.Notification{},
				ledger:       &ledger.Ledger{},
			}
        } else {
            fmt.Println("Single instance already created.")
        }
    } else {
        fmt.Println("Single instance already created.")
    }

	//the same singleton instance is returned every time by the GetInstance()
    return singleInstanceOfWalletService
    
}



func (w*WalletService)Deposit(accName ,codeByUser string) error {
	

	if w.account.CheckAccount(accName) && w.securityCode.CheckCode(codeByUser) { 
		amtDeposit := 0
		fmt.Println("Enter amount to be deposited")
		fmt.Scanln(&amtDeposit)
		w.wallet.Credit(amtDeposit)
		w.ledger.MakeEntry(accName, "deposited", amtDeposit)
		notificationMsg := w.notification.SendCreditNotification()
		fmt.Println(notificationMsg)
		return nil
	}

	err := errors.New("incorrect name or code")
	return err

}

func (w*WalletService)Withdraw(accName ,codeByUser string) error {
	
	if w.account.CheckAccount(accName) && w.securityCode.CheckCode(codeByUser) { //true
		amtWithdraw := 0
		withdraw:
		fmt.Println("Enter amount to be writhdrawn")
		fmt.Scanln(&amtWithdraw)
		err := w.wallet.Debit(amtWithdraw)
		if err != nil {
			fmt.Println(err)
			goto withdraw
		}
		w.ledger.MakeEntry(accName, "withdrawal", amtWithdraw)
		notificationMsg := w.notification.SendDebitNotification()
		fmt.Println(notificationMsg)
		return nil
	}

	err := errors.New("incorrect name or code")
	return err

}

func (w*WalletService)GetAllLedgerEntries(){
	w.ledger.GetAllEntries()
}