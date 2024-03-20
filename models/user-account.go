// File: account.go
package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAccount struct {
	gorm.Model    `json:"_"`
	ID            uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CustomerID    uuid.UUID `gorm:"foreignKey:CustomerID" json:"customer_id"`
	AccountTypeID uuid.UUID `gorm:"foreignKey:AccountTypeID" json:"account_type_id"`
	Balance       float64   `json:"balance"`
}
