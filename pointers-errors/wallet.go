package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin uint32

var ErrInsufficientBTC = errors.New("")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	w.balance += amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientBTC
	}
	w.balance -= amount
	return nil
}
