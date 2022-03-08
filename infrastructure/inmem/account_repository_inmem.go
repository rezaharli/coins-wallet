package inmem

import (
	accountEntity "behometest/domain/entity/account"
	"behometest/domain/repository"
)

type accountRepository struct {
	accounts map[accountEntity.AccountID]*accountEntity.Account
}

func (r *accountRepository) Get(id accountEntity.AccountID) (*accountEntity.Account, error) {
	if account, ok := r.accounts[id]; ok {
		return account, nil
	}
	return nil, accountEntity.ErrUnknown
}

func (r *accountRepository) GetAll() ([]accountEntity.Account, error) {
	accounts := make([]accountEntity.Account, 0, len(r.accounts))
	for _, val := range r.accounts {
		accounts = append(accounts, *val)
	}
	return accounts, nil
}

func (r *accountRepository) Create(account accountEntity.Account) error {
	if _, ok := r.accounts[account.ID]; ok {
		return accountEntity.ErrAlreadyExist
	}

	r.accounts[account.ID] = &account
	return nil
}

func NewAccountRepository() repository.AccountRepository {
	r := &accountRepository{
		accounts: make(map[accountEntity.AccountID]*accountEntity.Account),
	}
	r.accounts[`bob123`] = &accountEntity.Account{ID: `bob123`, Balance: 100.00, Currency: `USD`}
	r.accounts[`alice456`] = &accountEntity.Account{ID: `alice456`, Balance: 0.01, Currency: `USD`}
	r.accounts[`john789`] = &accountEntity.Account{ID: `john789`, Balance: 10, Currency: `AUD`}
	return r
}
