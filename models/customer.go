// File: customer.go
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model `json:"_"`
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	DOB        time.Time `json:"dob"`
	Address    string    `json:"address"`
	UserID     uuid.UUID `gorm:"foreignKey:UserID" json:"user_id"`
}
