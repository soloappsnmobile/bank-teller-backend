// File: teller_credentials.go
package models

import "gorm.io/gorm"

type TellerCredentials struct {
	gorm.Model `json:"_"`
	TellerID   int    `gorm:"foreignKey:TellerID" json:"teller_id"`
	Password   string `json:"password"`
}
