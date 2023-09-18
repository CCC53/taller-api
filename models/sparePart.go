package models

import (
	"taller-api/enums"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SparePart struct {
	ID           uuid.UUID             `gorm:"type:uuid;primary_key;" json:"id"`
	Name         string                `gorm:"not null" json:"name"`
	Disponible   int                   `gorm:"not null" json:"disponible"`
	Price        float64               `gorm:"not null" json:"price"`
	Supplier     string                `gorm:"not null" json:"supplier"`
	PurchaseDate time.Time             `gorm:"not null" json:"purchaseDate"`
	Type         enums.ValidSpareParts `gorm:"not null" json:"type"`
	ServiceID    *uuid.UUID            `json:"serviceID,omitempty"`
}

func (sparePart *SparePart) BeforeCreate(tx *gorm.DB) (err error) {
	sparePart.ID = uuid.New()
	return
}
