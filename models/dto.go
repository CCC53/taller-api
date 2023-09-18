package models

type LoginDTO struct {
	Email    string
	Password string
}

type AssignEmployeeDTO struct {
	EmployeeID string `json:"employeeID"`
}

type AssignSparePartDTO struct {
	SparePartID   string `json:"sparePartID"`
	QuantityToUse int    `json:"quantityToUse"`
}
