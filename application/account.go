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

func (s *account) Get(accountid accountEntity.AccountID) (*accountEntity.Account, error) {
	return s.accountRepository.Get(accountid)
}

func (s *account) GetAll() ([]accountEntity.Account, error) {
	return s.accountRepository.GetAll()
}

func (s *account) Create(account accountEntity.Account) error {
	return s.accountRepository.Create(account)
}
