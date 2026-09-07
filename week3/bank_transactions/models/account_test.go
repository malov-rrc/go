package models

import (
	"errors"
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
