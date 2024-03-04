// File: teller.go
package models

import "gorm.io/gorm"

type Teller struct {
	gorm.Model `json:"_"`
	TellerID   int    `gorm:"primaryKey" json:"teller_id"`
	Username   string `json:"username"`
}
