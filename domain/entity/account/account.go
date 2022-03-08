package account

import (
	"errors"
)

type AccountID string

type Account struct {
	ID       AccountID `json:"id" gorm:"primaryKey"`
	Balance  float32   `json:"balance"`
	Currency string    `json:"currency"`
}

var ErrUnknown = errors.New("unknown account")
var ErrAlreadyExist = errors.New("account already exist")
var ErrInsufficientBalance = errors.New("insufficient balance")
var ErrCurrencyMismatch = errors.New("currency mismatch")
