// File: account.go
package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model  `json:"_"`
	AccountID   int     `gorm:"primaryKey" json:"account_id"`
	CustomerID  int     `gorm:"foreignKey:CustomerID" json:"customer_id"`
	AccountType string  `json:"account_type"`
	Balance     float64 `json:"balance"`
}
