// File: customer.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model `json:"_"`
	CustomerID int       `gorm:"primaryKey" json:"customer_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	DOB        time.Time `json:"dob"`
	Address    string    `json:"address"`
	Phone      string    `json:"phone"`
}
