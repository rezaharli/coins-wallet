package repository

import (
	"coins-wallet/domain/entity/account"
	"coins-wallet/domain/entity/payment"
)

type PaymentRepository interface {
	GetAll() ([]payment.Payment, error)
	Save(from, to *account.Account, payments *[]payment.Payment) error
}
