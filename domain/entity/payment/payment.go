package payment

import (
	"errors"

	"coins-wallet/domain/entity/account"
)

type PaymentDirection string

type Payment struct {
	Account     account.AccountID `json:"account"`
	Amount      float32           `json:"amount"`
	ToAccount   account.AccountID `json:"to_account,omitempty"`
	FromAccount account.AccountID `json:"from_account,omitempty"`
	Direction   PaymentDirection  `json:"direction"`
}

var ErrUnknown = errors.New("unknown payment data")
var ErrAmountZero = errors.New("amount should be greater than 0")
