package application

import (
	accountEntity "coins-wallet/domain/entity/account"
	paymentEntity "coins-wallet/domain/entity/payment"
	"coins-wallet/domain/repository"
)

type Payment interface {
	GetAll() ([]paymentEntity.Payment, error)
	Save(sender, receiver *accountEntity.Account, payments *[]paymentEntity.Payment) error
}

var _ Payment = &payment{}

type payment struct {
	paymentRepository repository.PaymentRepository
}

// GetAll returns all payments
func (s *payment) GetAll() ([]paymentEntity.Payment, error) {
	return s.paymentRepository.GetAll()
}

// Save saves payments
func (s *payment) Save(sender, receiver *accountEntity.Account, payments *[]paymentEntity.Payment) error {
	return s.paymentRepository.Save(sender, receiver, payments)
}
