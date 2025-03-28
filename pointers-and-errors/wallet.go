package main

import (
	"errors"
	"fmt"
)

type Bitcoin float64
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount

}

var errInsufficientFunds = errors.New("oh no, not enough funds to withdraw")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}
