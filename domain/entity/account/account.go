package account

import (
	"errors"
)

// AccountID type of the account ID
type AccountID string

// Account represent entity of the account
type Account struct {
	ID       AccountID `json:"id" gorm:"primaryKey"`
	Balance  float32   `json:"balance"`
	Currency string    `json:"currency"`
}

// ErrUnknown unknown account error
var ErrUnknown = errors.New("unknown account")

// ErrAlreadyExist account already exist error
var ErrAlreadyExist = errors.New("account already exist")

// ErrInsufficientBalance insufficient balance error
var ErrInsufficientBalance = errors.New("insufficient balance")

// ErrCurrencyMismatch currency mismatch error
var ErrCurrencyMismatch = errors.New("currency mismatch")
