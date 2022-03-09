package wallet

import (
	"coins-wallet/application"
	accountEntity "coins-wallet/domain/entity/account"
	paymentEntity "coins-wallet/domain/entity/payment"
)

type Service interface {
	CreateAccount(accountEntity.Account) error
	GetAllAccount() ([]accountEntity.Account, error)
	GetAllPayment() ([]paymentEntity.Payment, error)
	Transfer(senderId, receiverId accountEntity.AccountID, amount float32) ([]paymentEntity.Payment, error)
}

type service struct {
	account application.Account
	payment application.Payment
}

func NewService(accountApp application.Account, paymentApp application.Payment) Service {
	return &service{
		account: accountApp,
		payment: paymentApp,
	}
}

func (s *service) CreateAccount(account accountEntity.Account) error {
	return s.account.Create(account)
}

func (s *service) GetAllAccount() ([]accountEntity.Account, error) {
	return s.account.GetAll()
}

func (s *service) GetAllPayment() ([]paymentEntity.Payment, error) {
	return s.payment.GetAll()
}

func (s *service) Transfer(senderId, receiverId accountEntity.AccountID, amount float32) ([]paymentEntity.Payment, error) {
	if amount <= 0 {
		return nil, paymentEntity.ErrAmountZero
	}

	sender, err := s.account.Get(senderId)
	if err != nil {
		return nil, err
	}

	if sender.Balance < amount {
		return nil, accountEntity.ErrInsufficientBalance
	}

	receiver, err := s.account.Get(receiverId)
	if err != nil {
		return nil, err
	}

	if sender.Currency != receiver.Currency {
		return nil, accountEntity.ErrCurrencyMismatch
	}

	sender.Balance -= amount
	receiver.Balance += amount

	payments := []paymentEntity.Payment{
		{
			Account:   sender.ID,
			Amount:    amount,
			ToAccount: receiver.ID,
			Direction: `outgoing`,
		},
		{
			Account:     receiver.ID,
			Amount:      amount,
			FromAccount: sender.ID,
			Direction:   `incoming`,
		},
	}

	err = s.payment.Save(sender, receiver, &payments)
	if err != nil {
		return nil, err
	}

	return payments, nil
}
