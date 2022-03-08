package application

import (
	accountEntity "behometest/domain/entity/account"
	paymentEntity "behometest/domain/entity/payment"
	"behometest/domain/repository"
)

type Payment interface {
	GetAll() ([]paymentEntity.Payment, error)
	Save(sender, receiver *accountEntity.Account, payments *[]paymentEntity.Payment) error
}

var _ Payment = &payment{}

type payment struct {
	paymentRepository repository.PaymentRepository
}

func (s *payment) GetAll() ([]paymentEntity.Payment, error) {
	return s.paymentRepository.GetAll()
}

func (s *payment) Save(sender, receiver *accountEntity.Account, payments *[]paymentEntity.Payment) error {
	return s.paymentRepository.Save(sender, receiver, payments)
}
