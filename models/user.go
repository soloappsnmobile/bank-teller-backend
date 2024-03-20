// File: user.go
package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"_"`
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	RoleID     uuid.UUID `gorm:"foreignKey:RoleID" json:"role_id"`
}
