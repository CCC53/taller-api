package models

import (
	"taller-api/enums"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID         uuid.UUID               `gorm:"type:uuid;primary_key;" json:"id"`
	Brand      string                  `gorm:"not null" json:"brand"`
	Model      string                  `gorm:"not null" json:"model"`
	Type       enums.ValidVehicleTypes `gorm:"not null" json:"type"`
	Year       int                     `gorm:"not null" json:"year"`
	Chassis    string                  `gorm:"not null" json:"chassis"`
	Motor      string                  `gorm:"not null" json:"motor"`
	Plate      string                  `gorm:"not null" json:"plate"`
	Owner      string                  `gorm:"not null" json:"owner"`
	EmailOwner string                  `gorm:"not null" json:"emailOwner"`
	Service    *Service                `gorm:"constraint:OnDelete:CASCADE" json:"service,omitempty"`
}

func (vehicle *Vehicle) BeforeCreate(tx *gorm.DB) (err error) {
	vehicle.ID = uuid.New()
	return
}
