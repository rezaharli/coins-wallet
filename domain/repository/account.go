package repository

import "coins-wallet/domain/entity/account"

// AccountRepository represent repository of account
// Expect implementation by the infrastructure layer
type AccountRepository interface {
	Get(id account.AccountID) (*account.Account, error)
	GetAll() ([]account.Account, error)
	Create(account.Account) error
}
