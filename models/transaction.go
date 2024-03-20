// File: transaction.go
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model        `json:"_"`
	ID                uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID            uuid.UUID `gorm:"foreignKey:UserID" json:"user_id"`
	AccountID         uuid.UUID `gorm:"foreignKey:AccountID" json:"account_id"`
	TransactionTypeID uuid.UUID `gorm:"foreignKey:TransactionTypeID" json:"transaction_type_id"`
	Amount            float64   `json:"amount"`
	TransactionDate   time.Time `json:"transaction_date"`
}
