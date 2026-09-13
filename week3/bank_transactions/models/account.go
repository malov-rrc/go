package models

import (
	"bank_transactions/utils"
	"fmt"
	"strings"
)

type Account struct {
	Balance float64
}

type rejectedEntry struct {
	Line int
	Err  error
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
	var rejectedTransactions []rejectedEntry
	var totalDeposited float64
	var totalWithdrawn float64
	var transactionsApplied int
	for i, line := range commandList {
		command, amount, err := utils.ParseLine(line)
		if err != nil {
			rejectedTransactions = append(rejectedTransactions, rejectedEntry{Line: i + 1, Err: err})
			continue
		}
		switch command {
		case string(Deposit):
			err := account.Deposit(amount)
			if err != nil {
				rejectedTransactions = append(rejectedTransactions, rejectedEntry{Line: i + 1, Err: err})
			} else {
				totalDeposited += amount
				transactionsApplied++
			}
		case string(Withdraw):
			err := account.Withdraw(amount)
			if err != nil {
				rejectedTransactions = append(rejectedTransactions, rejectedEntry{Line: i + 1, Err: err})
			} else {
				totalWithdrawn += amount
				transactionsApplied++
			}
		default:
			defaultError := &ValidationError{Field: "command", Msg: "unknown command: " + command}
			rejectedTransactions = append(rejectedTransactions, rejectedEntry{Line: i + 1, Err: defaultError})
		}
	}
	return account.getTransactionsStat(rejectedTransactions, totalDeposited, totalWithdrawn, transactionsApplied)
}

func (account *Account) getTransactionsStat(
	rejectedTransactions []rejectedEntry,
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
		for _, rejectedTransaction := range rejectedTransactions {
			fmt.Fprintf(&out, "%d: %s\n", rejectedTransaction.Line, rejectedTransaction.Err)
		}
	}
	return out.String()
}
