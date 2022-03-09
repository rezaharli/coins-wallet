package persistence

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"coins-wallet/domain/entity/account"
	"coins-wallet/domain/entity/payment"
	"coins-wallet/domain/repository"
)

type Repositories struct {
	db                *gorm.DB
	AccountRepository repository.AccountRepository
	PaymentRepository repository.PaymentRepository
}

// NewRepositories creates repository instance with db attached
func NewRepositories(DbHost, DbPort, DbUser, DbName, DbPassword string) (*Repositories, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repositories{
		db:                db,
		AccountRepository: NewAccountRepository(db),
		PaymentRepository: NewPaymentRepository(db),
	}, nil
}

// Automigrate migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&account.Account{}, &payment.Payment{})
}
