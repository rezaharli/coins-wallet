package repository

import (
	"behometest/domain/entity/account"
	"behometest/domain/entity/payment"
)

type PaymentRepository interface {
	GetAll() ([]payment.Payment, error)
	Save(from, to *account.Account, payments *[]payment.Payment) error
}
