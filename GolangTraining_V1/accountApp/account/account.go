package account

import (
	"errors"
	"math/rand"
	"time"
)

var allAccountSlice []*Account

type Account struct {
	AccNo   int
	Balance int
}

func New() *Account {
	var AccountNew = &Account{
		AccNo:   getNewId(),
		Balance: 0,
	}
	allAccountSlice = append(allAccountSlice, AccountNew)

	return AccountNew //this is a pointer
}

func getNewId() int {

	begin:
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	NewAccNo := rand.Intn(100) + 1
	for _, value := range allAccountSlice {
		if value.AccNo == NewAccNo {
			goto begin
		}
	}
	return NewAccNo

}

func (a *Account) Transfer(toAccNo, amt int) (string, error) {
	if (a.Balance - amt) < 1000 {
		err := errors.New("Money cannot be transferred since your account balance will become less than 1000")
		return "", err
	}
	success := false
	for _, accNoValue := range allAccountSlice {
		if accNoValue.AccNo == toAccNo {
			success = true
			a.Balance = a.Balance - amt
			accNoValue.Balance = accNoValue.Balance + amt
			break
		}
	}

	if success == false {
		err := errors.New("Account number not found")
		return "", err
	}

	return "Transfer done", nil

}

func (a *Account) Deposit(amt int) (string, error) {

	if amt < 0 {
		err := errors.New("Amount cannot be negative")
		return "", err
	}
	a.Balance = a.Balance + amt
	return "Amount has been deposited successfully", nil

}

func (a *Account) Withdraw(amt int) (string, error) {

	if a.Balance-amt < 1000 {
		err := errors.New("Amount cannot be withdraw as your account balance will become less than 1000 after withdrawal")
		return "", err
	}
	a.Balance = a.Balance - amt
	return "Amount has been withdrawn successfully", nil

}
