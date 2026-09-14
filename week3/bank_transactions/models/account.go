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

type TransactionResult struct {
	Rejected            []rejectedEntry
	TotalDeposited      float64
	TotalWithdrawn      float64
	TransactionsApplied int
	FinalBalance        float64
}

func (account *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return &ValidationError{"deposit", "amount must be positive"}
	}
	account.Balance += amount
	return nil
}

func (account *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return &ValidationError{"withdraw", "amount must be positive"}
	}
	if amount > account.Balance {
		return &InsufficientFundsError{Requested: amount, Available: account.Balance}
	}
	account.Balance -= amount
	return nil
}

// ProcessTransactions — вся бизнес-логика, никакого форматирования.
// Раньше это была первая половина MakeTransactions.
func (account *Account) ProcessTransactions(commandList []string) TransactionResult {
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

	return TransactionResult{
		Rejected:            rejectedTransactions,
		TotalDeposited:      totalDeposited,
		TotalWithdrawn:      totalWithdrawn,
		TransactionsApplied: transactionsApplied,
		FinalBalance:        account.Balance,
	}
}

// MakeTransactions — тонкая обёртка для main: вычислить + отформатировать.
// Старые вызовы в main.go не трогаем вообще.
func (account *Account) MakeTransactions(commandList []string) string {
	result := account.ProcessTransactions(commandList)
	return account.getTransactionsStat(result)
}

func (account *Account) getTransactionsStat(result TransactionResult) string {
	var out strings.Builder
	fmt.Fprintf(&out, "Final balance: %.2f\n", result.FinalBalance)
	fmt.Fprintf(&out, "Total deposited: %.2f\n", result.TotalDeposited)
	fmt.Fprintf(&out, "Total withdrawn: %.2f\n", result.TotalWithdrawn)
	fmt.Fprintf(&out, "Transactions applied: %d\n", result.TransactionsApplied)
	fmt.Fprintf(&out, "Transactions rejected: %d\n", len(result.Rejected))
	if len(result.Rejected) > 0 {
		fmt.Fprintln(&out, "Rejected details:")
		for _, r := range result.Rejected {
			fmt.Fprintf(&out, "%d: %s\n", r.Line, r.Err.Error())
		}
	}
	return out.String()
}
