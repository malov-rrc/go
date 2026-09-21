package models

type Command string

const (
	Withdraw Command = "WITHDRAW"
	Deposit  Command = "DEPOSIT"
)

func (s Command) String() string {
	switch s {
	case Withdraw:
		return "WITHDRAW"
	case Deposit:
		return "DEPOSIT"
	default:
		return "unknown"
	}
}
