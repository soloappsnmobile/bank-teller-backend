// File: transaction.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model      `json:"_"`
	TransactionID   int       `gorm:"primaryKey" json:"transaction_id"`
	TellerID        int       `gorm:"foreignKey:TellerID" json:"teller_id"`
	AccountID       int       `gorm:"foreignKey:AccountID" json:"account_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
}
