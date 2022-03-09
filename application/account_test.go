package application

import (
	"testing"

	accountEntity "coins-wallet/domain/entity/account"
	"coins-wallet/infrastructure/inmem"

	"github.com/stretchr/testify/assert"
)

func TestGetAccount_Success(t *testing.T) {
	var acc Account = &account{accountRepository: inmem.NewAccountRepository()}
	account, err := acc.Get(`bob123`)
	assert.Nil(t, err)
	assert.Equal(t, accountEntity.AccountID("bob123"), account.ID)
	assert.Equal(t, float32(100), account.Balance)
	assert.Equal(t, `USD`, account.Currency)
}

func TestGetAccount_NotFound(t *testing.T) {
	var acc Account = &account{accountRepository: inmem.NewAccountRepository()}
	account, err := acc.Get(`fake999`)
	assert.Error(t, accountEntity.ErrUnknown, err)
	assert.Nil(t, account)
}

func TestGetAllAccount_Success(t *testing.T) {
	var acc Account = &account{accountRepository: inmem.NewAccountRepository()}
	accounts, err := acc.GetAll()
	assert.Nil(t, err)
	assert.Len(t, accounts, 3)
}

func TestCreateAccount_Success(t *testing.T) {
	var acc Account = &account{accountRepository: inmem.NewAccountRepository()}
	err := acc.Create(accountEntity.Account{ID: `fake999`, Balance: 10, Currency: "AUD"})
	assert.Nil(t, err)

	accounts, err := acc.GetAll()
	assert.Nil(t, err)
	assert.Len(t, accounts, 4)
}

func TestCreateAccount_FailExist(t *testing.T) {
	var acc Account = &account{accountRepository: inmem.NewAccountRepository()}
	err := acc.Create(accountEntity.Account{ID: `bob123`, Balance: 10, Currency: "AUD"})
	assert.Error(t, accountEntity.ErrAlreadyExist, err)

	accounts, err := acc.GetAll()
	assert.Nil(t, err)
	assert.Len(t, accounts, 3)
}
