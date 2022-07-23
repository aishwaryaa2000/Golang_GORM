package service

import (
	"fmt"
	"wallet/components/account"
	"wallet/components/ledger"
	"wallet/components/securitycode"
	"wallet/components/wallet"
	"wallet/components/notification"
	"errors"

)


type WalletService struct {
    account      *account.Account
    wallet       *wallet.Wallet
    securityCode *securitycode.SecurityCode
    notification *notification.Notification
    ledger       *ledger.Ledger
}


func NewWalletService() *WalletService {
    fmt.Println("Starting create account")
    walletFacacde := &WalletService{
        account:      account.New(),
        securityCode: securitycode.New(),
        wallet:       wallet.New(),
        notification: &notification.Notification{},
        ledger:       &ledger.Ledger{},
    }
    fmt.Println("Account created")
    return walletFacacde
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