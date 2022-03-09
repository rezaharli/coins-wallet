package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"coins-wallet/domain/entity/account"
	"coins-wallet/domain/entity/payment"
	"coins-wallet/infrastructure/inmem"
)

func TestCreateAccount_Success(t *testing.T) {
	service := NewService(inmem.NewAccountRepository(), inmem.NewPaymentRepository())
	err := service.CreateAccount(account.Account{ID: `fake999`, Balance: 10, Currency: `AUD`})
	assert.Nil(t, err)

	accounts, err := service.GetAllAccount()
	assert.Nil(t, err)
	assert.Len(t, accounts, 4)
}

func TestGetAllAccounts_Success(t *testing.T) {
	service := NewService(inmem.NewAccountRepository(), inmem.NewPaymentRepository())
	accounts, err := service.GetAllAccount()
	assert.Nil(t, err)
	assert.Len(t, accounts, 3)
}

func TestGetAllPayments_Success(t *testing.T) {
	service := NewService(inmem.NewAccountRepository(), inmem.NewPaymentRepository())
	payments, err := service.GetAllPayment()
	assert.Nil(t, err)
	assert.Len(t, payments, 0)
}

func TestTransfer_Success(t *testing.T) {
	service := NewService(inmem.NewAccountRepository(), inmem.NewPaymentRepository())
	payments, err := service.Transfer(account.AccountID(`bob123`), account.AccountID(`alice456`), 100)
	assert.Nil(t, err)
	assert.Len(t, payments, 2)
}

func TestTransfer_ErrAmountZero(t *testing.T) {
	service := NewService(inmem.NewAccountRepository(), inmem.NewPaymentRepository())
	payments, err := service.Transfer(account.AccountID(`bob123`), account.AccountID(`alice456`), 0)
	assert.Error(t, payment.ErrAmountZero, err)
	assert.Len(t, payments, 0)
}

func TestTransfer_ErrUnknown(t *testing.T) {
	service := NewService(inmem.NewAccountRepository(), inmem.NewPaymentRepository())
	payments, err := service.Transfer(account.AccountID(`bob1234`), account.AccountID(`alice456`), 100)
	assert.Error(t, account.ErrUnknown, err)
	assert.Len(t, payments, 0)

	payments, err = service.Transfer(account.AccountID(`bob123`), account.AccountID(`alice4567`), 100)
	assert.Error(t, account.ErrUnknown, err)
	assert.Len(t, payments, 0)
}

func TestTransfer_ErrInsufficientBalance(t *testing.T) {
	service := NewService(inmem.NewAccountRepository(), inmem.NewPaymentRepository())
	payments, err := service.Transfer(account.AccountID(`bob123`), account.AccountID(`alice456`), 1000)
	assert.Error(t, account.ErrInsufficientBalance, err)
	assert.Len(t, payments, 0)
}

func TestTransfer_ErrCurrencyMismatch(t *testing.T) {
	service := NewService(inmem.NewAccountRepository(), inmem.NewPaymentRepository())
	payments, err := service.Transfer(account.AccountID(`bob123`), account.AccountID(`john789`), 100)
	assert.Error(t, account.ErrCurrencyMismatch, err)
	assert.Len(t, payments, 0)
}
