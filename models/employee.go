package models

import (
	"taller-api/enums"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID        uuid.UUID        `gorm:"type:uuid;primary_key;"`
	Name      string           `gorm:"not null"`
	Email     string           `gorm:"not null; unique"`
	Password  string           `gorm:"not null size:5"`
	Role      enums.ValidRoles `gorm:"not null"`
	ServiceID *uuid.UUID       `json:"serviceID,omitempty"`
}

type EmployeeResponse struct {
	ID        uuid.UUID        `json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Role      enums.ValidRoles `json:"role"`
	ServiceID *uuid.UUID       `json:"serviceID,omitempty"`
}

func (employee *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	employee.ID = uuid.New()
	return
}
