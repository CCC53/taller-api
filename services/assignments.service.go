package services

import (
	"taller-api/db"
	"taller-api/enums"
	"taller-api/models"

	"github.com/google/uuid"
)

func mapEmployeesSelect(employees []models.Employee) []models.SelectResponse {
	employeesSelect := []models.SelectResponse{}
	for _, employee := range employees {
		employeeMapped := models.SelectResponse{
			Value: employee.ID,
			Label: employee.Name,
		}
		employeesSelect = append(employeesSelect, employeeMapped)
	}
	return employeesSelect
}

func mapSaprePartsSelect(spareParts []models.SparePart) []models.SelectResponse {
	sparePartSelect := []models.SelectResponse{}
	for _, sparePart := range spareParts {
		sparePartMapped := models.SelectResponse{
			Value: sparePart.ID,
			Label: sparePart.Name,
		}
		sparePartSelect = append(sparePartSelect, sparePartMapped)
	}
	return sparePartSelect
}

func ListMechanicsAviable() []models.SelectResponse {
	var employees []models.Employee
	db.DB.Find(&employees, "role = ? AND service_id IS NULL", enums.Mechanic)
	return mapEmployeesSelect(employees)
}

func ListSaparePartsAviable() []models.SelectResponse {
	var spareParts []models.SparePart
	db.DB.Find(&spareParts, "disponible > 0")
	return mapSaprePartsSelect(spareParts)
}

func AssignEmployeeToService(id string, employeeID string) (models.EmployeeResponse, error) {
	employee, err := GetEmployeeByID(employeeID)
	var idParsed = uuid.MustParse(id)
	if err != nil {
		return models.EmployeeResponse{}, err
	}
	employee.ServiceID = &idParsed
	db.DB.Save(&employee)
	return MappingEmployee(employee), nil
}

func AssignSparePartToService(id string, data models.AssignSparePartDTO) (models.SparePart, error) {
	sparePart, err := GetSparePartByID(data.SparePartID)
	var idParsed = uuid.MustParse(id)
	if err != nil {
		return models.SparePart{}, err
	}
	sparePart.ServiceID = &idParsed
	sparePart.Disponible = sparePart.Disponible - data.QuantityToUse
	db.DB.Save(&sparePart)
	return sparePart, nil
}

func RemoveEmployeeFromService(id string) (bool, error) {
	employee, err := GetEmployeeByID(id)
	if err != nil {
		return false, err
	}
	employee.ServiceID = nil
	db.DB.Save(&employee)
	return true, nil
}

func RemoveSparePartFromService(id string) (bool, error) {
	sparePart, err := GetSparePartByID(id)
	if err != nil {
		return false, err
	}
	sparePart.ServiceID = nil
	db.DB.Save(&sparePart)
	return true, nil
}
