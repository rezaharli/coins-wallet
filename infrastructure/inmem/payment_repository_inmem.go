package inmem

import (
	"coins-wallet/domain/entity/account"
	"coins-wallet/domain/entity/payment"
	paymentEntity "coins-wallet/domain/entity/payment"
	"coins-wallet/domain/repository"
)

type paymentRepository struct {
	payments []*paymentEntity.Payment
}

// GetAll mimics GetAll from payment_repository
func (r *paymentRepository) GetAll() ([]paymentEntity.Payment, error) {
	payments := make([]paymentEntity.Payment, 0, len(r.payments))
	for _, val := range r.payments {
		payments = append(payments, *val)
	}
	return payments, nil
}

// Save mimics Save from payment_repository
func (r *paymentRepository) Save(from, to *account.Account, payments *[]payment.Payment) error {
	for _, payment := range *payments {
		r.payments = append(r.payments, &payment)
	}
	return nil
}

// NewPaymentRepository creates new inmem repository for payment
func NewPaymentRepository() repository.PaymentRepository {
	r := &paymentRepository{
		payments: make([]*paymentEntity.Payment, 0),
	}

	return r
}
