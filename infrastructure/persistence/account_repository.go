package persistence

import (
	"strings"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"

	accountEntity "coins-wallet/domain/entity/account"
	"coins-wallet/domain/repository"
)

type AccountRepositoryImpl struct {
	Conn *gorm.DB
}

// NewAccountRepository creates new repository for account
func NewAccountRepository(conn *gorm.DB) repository.AccountRepository {
	return &AccountRepositoryImpl{Conn: conn}
}

// Get gets account data by id from db
func (r *AccountRepositoryImpl) Get(id accountEntity.AccountID) (*accountEntity.Account, error) {
	account := &accountEntity.Account{}
	if err := r.Conn.First(&account, id).Error; err != nil {
		return nil, err
	}
	return account, nil
}

// GetAll gets all account datas from db
func (r *AccountRepositoryImpl) GetAll() ([]accountEntity.Account, error) {
	accounts := []accountEntity.Account{}
	if err := r.Conn.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

// Create new account data to db
func (r *AccountRepositoryImpl) Create(account accountEntity.Account) error {
	if err := r.Conn.Create(&account).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return accountEntity.ErrAlreadyExist
		}
		return err
	}
	return nil
}
