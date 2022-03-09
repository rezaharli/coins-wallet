package application

import (
	accountEntity "coins-wallet/domain/entity/account"
	"coins-wallet/domain/repository"
)

type Account interface {
	Get(accountid accountEntity.AccountID) (*accountEntity.Account, error)
	GetAll() ([]accountEntity.Account, error)
	Create(accountEntity.Account) error
}

var _ Account = &account{}

type account struct {
	accountRepository repository.AccountRepository
}

// Get returns account by id
func (s *account) Get(accountid accountEntity.AccountID) (*accountEntity.Account, error) {
	return s.accountRepository.Get(accountid)
}

// GetAll returns all accounts
func (s *account) GetAll() ([]accountEntity.Account, error) {
	return s.accountRepository.GetAll()
}

// Create creates new accounts
func (s *account) Create(account accountEntity.Account) error {
	return s.accountRepository.Create(account)
}
