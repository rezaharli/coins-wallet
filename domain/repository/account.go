package repository

import "behometest/domain/entity/account"

type AccountRepository interface {
	Get(id account.AccountID) (*account.Account, error)
	GetAll() ([]account.Account, error)
	Create(account.Account) error
}
