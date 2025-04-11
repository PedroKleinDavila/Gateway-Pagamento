package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID             string `gorm:"primaryKey"`
	AccountID      string `gorm:"foreignKey:AccountID"`
	Status         Status
	Description    string
	Amount         float64
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number         string
	CVV            string
	ExpiryMonth    int
	ExpiryYear     int
	CardHolderName string
}

func NewInvoice(accountID string, amount float64, description string, paymentType string, card CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	lastDigits := card.Number[len(card.Number)-4:]

	return &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Status:         StatusPending,
		Description:    description,
		Amount:         amount,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (i *Invoice) Process() error {
	if i.Amount > 10000 {
		return nil
	}
	randomSource := rand.New(rand.NewSource(time.Now().Unix()))
	if randomSource.Float64() <= 0.7 {
		i.Status = StatusApproved
	} else {
		i.Status = StatusRejected
	}
	return nil
}

func (i *Invoice) UpdateStatus(status Status) error {
	if status != StatusPending {
		return ErrInvalidStatus
	}
	i.Status = status
	i.UpdatedAt = time.Now()
	return nil
}
