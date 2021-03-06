package application

import (
	"testing"

	accountEntity "coins-wallet/domain/entity/account"
	paymentEntity "coins-wallet/domain/entity/payment"
	"coins-wallet/infrastructure/inmem"

	"github.com/stretchr/testify/assert"
)

func TestSavePayment_Success(t *testing.T) {
	var pymnt Payment = &payment{paymentRepository: inmem.NewPaymentRepository()}

	paymentsIn := &[]paymentEntity.Payment{{Account: `bob123`, Amount: 100, ToAccount: `alice456`, Direction: `outgoing`}}
	err := pymnt.Save(&accountEntity.Account{}, &accountEntity.Account{}, paymentsIn)
	assert.Nil(t, err)

	payments, err := pymnt.GetAll()
	assert.Nil(t, err)
	assert.Len(t, payments, 1)
}
