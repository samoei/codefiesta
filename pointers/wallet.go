package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%dBTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) GetBalance() Bitcoin {
	return w.balance
}
