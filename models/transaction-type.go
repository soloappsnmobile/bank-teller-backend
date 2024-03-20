package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionType struct {
	gorm.Model `json:"_"`
	Name       string    `json:"name"` // Name can be "Deposit" or "Withdrawal"
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
}
