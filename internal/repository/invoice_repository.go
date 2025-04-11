package repository

import (
	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
	"gorm.io/gorm"
)

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) Save(invoice *domain.Invoice) error {
	return r.db.Create(invoice).Error
}

func (r *InvoiceRepository) FindByID(id string) (*domain.Invoice, error) {
	var invoice domain.Invoice
	result := r.db.First(&invoice, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, domain.ErrInvoiceNotFound
		}
		return nil, result.Error
	}
	return &invoice, nil
}

func (r *InvoiceRepository) FindByAccountID(accountID string) ([]*domain.Invoice, error) {
	var invoices []*domain.Invoice
	result := r.db.Where("account_id = ?", accountID).Find(&invoices)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, domain.ErrInvoiceNotFound
		}
		return nil, result.Error
	}
	return invoices, nil
}

func (r *InvoiceRepository) UpdateStatus(invoice *domain.Invoice) error {
	result := r.db.Model(&domain.Invoice{}).Where("id = ?", invoice.ID).Updates(map[string]interface{}{
		"status":     invoice.Status,
		"updated_at": invoice.UpdatedAt,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrInvoiceNotFound
	}
	return nil
}
