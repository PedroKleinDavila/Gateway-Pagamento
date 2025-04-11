package repository

import (
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	result := r.db.Where("api_key = ?", apiKey).First(&account)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, domain.ErrAccountNotFound
		}
		return nil, result.Error
	}
	return &account, nil
}

func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	var account domain.Account
	result := r.db.First(&account, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, domain.ErrAccountNotFound
		}
		return nil, result.Error
	}
	return &account, nil
}

func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var dbAccount domain.Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&dbAccount, "id = ?", account.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return domain.ErrAccountNotFound
			}
			return err
		}

		dbAccount.Balance += account.Balance
		dbAccount.UpdatedAt = time.Now()

		return tx.Save(&dbAccount).Error
	})
}
