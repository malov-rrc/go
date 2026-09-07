package models

import "errors"

type Account struct {
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	a.Balance += amount
	return nil
}
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	if amount > a.Balance {
		return &InsufficientFundsError{Requested: amount, Available: a.Balance}
	}
	a.Balance -= amount
	return nil
}
