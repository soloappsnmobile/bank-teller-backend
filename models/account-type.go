package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountType struct {
	gorm.Model `json:"_"`
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       string    `json:"name"` // Name can be "Savings" or "Current"
}
