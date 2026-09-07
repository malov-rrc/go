package models

import (
	"bank_transactions/utils"
	"errors"
	"fmt"
	"strings"
)

type Account struct {
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return &ValidationError{"deposit", "amount must be positive"}
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return &ValidationError{"withdraw", "amount must be positive"}
	}
	if amount > a.Balance {
		return &InsufficientFundsError{Requested: amount, Available: a.Balance}
	}
	a.Balance -= amount
	return nil
}

func (account *Account) MakeTransactions(commandList []string) string {
	rejectedTransactions := make(map[int]error)
	var totalDeposited float64
	var totalWithdrawn float64
	var transactionsApplied int
	for i, line := range commandList {
		command, amount, err := utils.ParseLine(line)
		if err != nil {
			rejectedTransactions[i+1] = err
		}
		switch command {
		case string(Deposit):
			err := account.Deposit(amount)
			if err != nil {
				rejectedTransactions[i+1] = err
			} else {
				totalDeposited += amount
				transactionsApplied++
			}
		case string(Withdraw):
			err := account.Withdraw(amount)
			if err != nil {
				rejectedTransactions[i+1] = err
			} else {
				totalWithdrawn += amount
				transactionsApplied++
			}
		default:
			rejectedTransactions[i+1] = errors.New("unknown command: " + command)
		}
	}
	return account.getTransactionsStat(rejectedTransactions, totalDeposited, totalWithdrawn, transactionsApplied)
}

func (account *Account) getTransactionsStat(
	rejectedTransactions map[int]error,
	deposited float64,
	withdrawn float64,
	transactionsApplied int,
) string {
	var out strings.Builder
	fmt.Fprintf(&out, "Final balance: %.2f\n", account.Balance)
	fmt.Fprintf(&out, "Total deposited: %.2f\n", deposited)
	fmt.Fprintf(&out, "Total withdrawn: %.2f\n", withdrawn)
	fmt.Fprintf(&out, "Transactions applied: %d\n", transactionsApplied)
	fmt.Fprintf(&out, "Transactions rejected: %d\n", len(rejectedTransactions))
	if len(rejectedTransactions) > 0 {
		fmt.Fprintln(&out, "Rejected details:")
		for i, err := range rejectedTransactions {
			fmt.Fprintf(&out, "%d: %s\n", i, err)
		}
	}
	return out.String()
}
