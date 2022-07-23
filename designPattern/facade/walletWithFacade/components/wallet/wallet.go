package wallet

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance int
}

func New() *Wallet {
	var a = &Wallet{
		balance: 0,
	}
	return a
}

func (w *Wallet) Credit(amt int) {
	w.balance =w.balance+amt
	fmt.Println("Amount deposited successfully and balance is : ",w.balance)
}

func (w *Wallet) Debit(amt int) error {
	if w.balance-amt < 0 {
		err:= errors.New("cannot debit as balance will become less than 0")
		return err
	}

	w.balance=w.balance-amt
	fmt.Println("Withdrawn successfully.Balance is : ",w.balance)
	return nil
}