package repository

import (
	"coins-wallet/domain/entity/account"
	"coins-wallet/domain/entity/payment"
)

// PaymentRepository represent repository of payment
// Expect implementation by the infrastructure layer
type PaymentRepository interface {
	GetAll() ([]payment.Payment, error)
	Save(from, to *account.Account, payments *[]payment.Payment) error
}
