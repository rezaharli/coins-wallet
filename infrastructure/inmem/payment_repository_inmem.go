package inmem

import (
	"behometest/domain/entity/account"
	"behometest/domain/entity/payment"
	paymentEntity "behometest/domain/entity/payment"
	"behometest/domain/repository"
)

type paymentRepository struct {
	payments []*paymentEntity.Payment
}

func (r *paymentRepository) GetAll() ([]paymentEntity.Payment, error) {
	payments := make([]paymentEntity.Payment, 0, len(r.payments))
	for _, val := range r.payments {
		payments = append(payments, *val)
	}
	return payments, nil
}

func (r *paymentRepository) Save(from, to *account.Account, payments *[]payment.Payment) error {
	for _, payment := range *payments {
		r.payments = append(r.payments, &payment)
	}
	return nil
}

func NewPaymentRepository() repository.PaymentRepository {
	r := &paymentRepository{
		payments: make([]*paymentEntity.Payment, 0),
	}

	return r
}
