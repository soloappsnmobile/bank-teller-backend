package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model `json:"_"`
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       string    `json:"name"` // Name can be "Customer", "Teller", or "Admin"
}

func GetRoleByName(db *gorm.DB, name string) (*Role, error) {
	var role Role
	if err := db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
