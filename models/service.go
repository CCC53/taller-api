package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	ID         uuid.UUID    `gorm:"type:uuid;primary_key;" json:"id"`
	StartDate  time.Time    `gorm:"not null" json:"startDate"`
	EndDate    time.Time    `gorm:"not null" json:"endDate"`
	Issue      string       `gorm:"not null" json:"issue"`
	VehicleID  *uuid.UUID   `json:"vehicleID"`
	Employees  *[]Employee  `gorm:"constraint:OnDelete:SET NULL" json:"employees,omitempty"`
	SpareParts *[]SparePart `gorm:"constraint:OnDelete:SET NULL" json:"spareParts,omitempty"`
}

type ServiceResponse struct {
	ID         uuid.UUID          `json:"id"`
	StartDate  time.Time          `json:"startDate"`
	EndDate    time.Time          `json:"endDate"`
	Issue      string             `json:"issue"`
	VehicleID  uuid.UUID          `json:"vehicleID"`
	Employees  []EmployeeResponse `json:"employees"`
	SpareParts []SparePart        `json:"spareParts"`
}

func (service *Service) BeforeCreate(tx *gorm.DB) (err error) {
	service.ID = uuid.New()
	return
}
