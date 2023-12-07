package services

import (
	"strconv"
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

func mapVehiclesSelect(vehicles []models.Vehicle) []models.SelectResponse {
	vehiclesSelect := []models.SelectResponse{}
	for _, vehicle := range vehicles {
		label := vehicle.Brand + "-" + vehicle.Model + "-" + strconv.Itoa(vehicle.Year)
		vehicleMapped := models.SelectResponse{
			Value: vehicle.ID,
			Label: label,
		}
		vehiclesSelect = append(vehiclesSelect, vehicleMapped)
	}
	return vehiclesSelect
}

func ListItemsAviable(table string) []models.SelectResponse {
	var data []models.SelectResponse
	switch table {
	case "employees":
		var employees []models.Employee
		db.DB.Find(&employees, "role = ? AND service_id IS NULL", enums.Mechanic)
		data = mapEmployeesSelect(employees)
	case "vehicles":
		var vehicles []models.Vehicle
		db.DB.Find(&vehicles)
		data = mapVehiclesSelect(vehicles)
	case "spare-parts":
		var spareParts []models.SparePart
		db.DB.Find(&spareParts, "disponible > 0")
		data = mapSaprePartsSelect(spareParts)
	}
	return data
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

func RemoveItemFromService(table string, id string) (bool, error) {
	if table == "employee" {
		employee, err := GetEmployeeByID(id)
		if err != nil {
			return false, err
		}
		employee.ServiceID = nil
		db.DB.Save(&employee)
		return true, nil
	}
	sparePart, err := GetSparePartByID(id)
	if err != nil {
		return false, err
	}
	sparePart.ServiceID = nil
	db.DB.Save(&sparePart)
	return true, nil
}
