package payment

import (
	"errors"

	"coins-wallet/domain/entity/account"
)

// PaymentDirection type of payment's direction of transfer
type PaymentDirection string

// Payment represent entity of the account
type Payment struct {
	Account     account.AccountID `json:"account"`
	Amount      float32           `json:"amount"`
	ToAccount   account.AccountID `json:"to_account,omitempty"`
	FromAccount account.AccountID `json:"from_account,omitempty"`
	Direction   PaymentDirection  `json:"direction"`
}

// ErrUnknown unknown payment data error
var ErrUnknown = errors.New("unknown payment data")

// ErrAmountZero amount should be greater than 0 error
var ErrAmountZero = errors.New("amount should be greater than 0")
