package persistence

import (
	"sync"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"

	accountEntity "coins-wallet/domain/entity/account"
	paymentEntity "coins-wallet/domain/entity/payment"
	"coins-wallet/domain/repository"
)

type PaymentRepositoryImpl struct {
	mtx  sync.Mutex
	Conn *gorm.DB
}

func NewPaymentRepository(conn *gorm.DB) repository.PaymentRepository {
	return &PaymentRepositoryImpl{Conn: conn}
}

func (r *PaymentRepositoryImpl) GetAll() ([]paymentEntity.Payment, error) {
	payment := []paymentEntity.Payment{}
	if err := r.Conn.Find(&payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepositoryImpl) Save(sender, receiver *accountEntity.Account, payments *[]paymentEntity.Payment) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	tx := r.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(&sender).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(&receiver).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := r.Conn.Create(&payments).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
