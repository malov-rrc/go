package models

import "fmt"

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string { return e.Msg }

type InsufficientFundsError struct {
	Requested float64
	Available float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds: requested %.2f, available %.2f", e.Requested, e.Available)
}
