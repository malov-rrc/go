package models

import (
	"errors"
	"fmt"
	"testing"
)

func TestAccount_Deposit(t *testing.T) {
	tests := []struct {
		name            string
		account         Account
		amount          float64
		expectedBalance float64
		checkError      func(err error) bool
	}{
		{
			name:            "normal deposit",
			account:         Account{Balance: 0},
			amount:          10,
			expectedBalance: 10,
			checkError:      func(err error) bool { return err == nil },
		},
		{
			name:            "negative deposit",
			account:         Account{Balance: 100},
			amount:          -10,
			expectedBalance: 100,
			checkError: func(err error) bool {
				var ve *ValidationError
				if !errors.As(err, &ve) {
					return false
				}
				return ve.Field == "deposit" && ve.Msg == "amount must be positive"
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			initBalance := tc.account.Balance
			err := tc.account.Deposit(tc.amount)

			if !tc.checkError(err) {
				t.Errorf("unexpected error: %v", err)
			}

			if tc.account.Balance != tc.expectedBalance {
				t.Errorf("balance = %.2f, expected %.2f (initial: %.2f, amount: %.2f)",
					tc.account.Balance, tc.expectedBalance, initBalance, tc.amount)
			}
		})
	}
}

func TestAccount_Withdraw(t *testing.T) {
	tests := []struct {
		name            string
		account         Account
		amount          float64
		expectedBalance float64
		checkError      func(err error) bool
	}{
		{
			name:            "normal withdraw",
			account:         Account{Balance: 100},
			amount:          30,
			expectedBalance: 70,
			checkError:      func(err error) bool { return err == nil },
		},
		{
			name:            "withdraw with negative amount",
			account:         Account{Balance: 100},
			amount:          -10,
			expectedBalance: 100,
			checkError: func(err error) bool {
				var ve *ValidationError
				if !errors.As(err, &ve) {
					return false
				}
				return ve.Field == "withdraw" && ve.Msg == "amount must be positive"
			},
		},
		{
			name:            "insufficient funds",
			account:         Account{Balance: 50},
			amount:          100,
			expectedBalance: 50,
			checkError: func(err error) bool {
				var ie *InsufficientFundsError
				if !errors.As(err, &ie) {
					return false
				}
				return ie.Requested == 100 && ie.Available == 50
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			initBalance := tc.account.Balance
			err := tc.account.Withdraw(tc.amount)

			if !tc.checkError(err) {
				t.Errorf("unexpected error: %v", err)
			}

			if tc.account.Balance != tc.expectedBalance {
				t.Errorf("balance = %.2f, expected %.2f (initial: %.2f, amount: %.2f)",
					tc.account.Balance, tc.expectedBalance, initBalance, tc.amount)
			}
		})
	}
}

func TestProcessTransactions_UnknownCommand_ReturnsValidationError(t *testing.T) {
	account := &Account{}
	result := account.ProcessTransactions([]string{"INVALID_COMMAND 10"})

	if len(result.Rejected) != 1 {
		t.Fatalf("expected 1 rejected transaction, got %d", len(result.Rejected))
	}

	var validationErr *ValidationError
	if !errors.As(result.Rejected[0].Err, &validationErr) {
		t.Fatalf("expected *ValidationError, got %T", result.Rejected[0].Err)
	}
	if validationErr.Field != "command" {
		t.Errorf("expected Field=command, got %q", validationErr.Field)
	}
}

func TestProcessTransactions_WithdrawTooMuch_ReturnsInsufficientFundsError(t *testing.T) {
	account := &Account{Balance: 100}
	result := account.ProcessTransactions([]string{"WITHDRAW 200"})

	var fundsErr *InsufficientFundsError
	if len(result.Rejected) != 1 {
		t.Fatalf("expected 1 rejected transaction, got %d", len(result.Rejected))
	}
	if !errors.As(result.Rejected[0].Err, &fundsErr) {
		t.Fatalf("expected *InsufficientFundsError, got %T", result.Rejected[0].Err)
	}
}

func TestMakeTransactions_FullIntegration(t *testing.T) {
	account := &Account{}
	result := account.MakeTransactions([]string{"DEPOSIT 100", "INVALID 10", "WITHDRAW 50"})
	fmt.Println(result)
	expectedResult := "Final balance: 50.00\n" +
		"Total deposited: 100.00\n" +
		"Total withdrawn: 50.00\n" +
		"Transactions applied: 2\n" +
		"Transactions rejected: 1\n" +
		"Rejected details:\n" +
		"2: unknown command: INVALID\n"
	if result != expectedResult {
		t.Errorf(
			"ERROR\n"+
				"\n============= expected output =============\n"+
				"%s"+
				"\n============= got output =============\n"+
				"%s",
			expectedResult, result,
		)
	}
}
