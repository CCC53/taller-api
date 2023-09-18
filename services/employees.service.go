package services

import (
	"taller-api/db"
	"taller-api/models"

	"gorm.io/gorm"
)

func MappingEmployees(employees []models.Employee) []models.EmployeeResponse {
	employeesMapped := []models.EmployeeResponse{}
	for _, employee := range employees {
		employeeMapped := models.EmployeeResponse{
			ID:        employee.ID,
			Name:      employee.Name,
			Email:     employee.Email,
			Role:      employee.Role,
			ServiceID: employee.ServiceID,
		}
		employeesMapped = append(employeesMapped, employeeMapped)
	}
	return employeesMapped
}

func MappingEmployee(employee models.Employee) models.EmployeeResponse {
	employeeMapped := models.EmployeeResponse{
		ID:        employee.ID,
		Name:      employee.Name,
		Email:     employee.Email,
		Role:      employee.Role,
		ServiceID: employee.ServiceID,
	}
	return employeeMapped
}

func ListEmployees(pageSize int, page int, token string) ([]models.EmployeeResponse, int64) {
	var employees []models.Employee
	var totalCount int64
	var id = GetMyID(token)
	db.DB.Scopes(Paginate(pageSize, page)).Find(&employees, "id != ? ORDER BY role", id)
	db.DB.Model(&employees).Where("id != ?", id).Count(&totalCount)
	return MappingEmployees(employees), totalCount
}

func GetEmployeeByID(id string) (models.Employee, error) {
	var employee models.Employee
	employeeDB := db.DB.First(&employee, "id = ?", id)
	if employeeDB.Error != nil {
		return models.Employee{}, employeeDB.Error
	}
	return employee, nil
}

func CreateEmployee(formData models.Employee) (models.EmployeeResponse, error) {
	employee := formData
	paswword, _ := HashPassword("luneta")
	employee.Password = paswword
	var employeeDB *gorm.DB = db.DB.Create(&employee)
	if employeeDB.Error != nil {
		return models.EmployeeResponse{}, employeeDB.Error
	}
	return MappingEmployee(employee), nil
}

func UpdateEmployee(id string, formData models.Employee) (models.EmployeeResponse, error) {
	employeeFound, err := GetEmployeeByID(id)
	if err != nil {
		return models.EmployeeResponse{}, err
	}
	employeeFound.Name = formData.Name
	employeeFound.Email = formData.Email
	employeeFound.Role = formData.Role
	db.DB.Save(&employeeFound)
	return MappingEmployee(employeeFound), nil
}

func DeleteEmployee(id string) (bool, error) {
	employee, err := GetEmployeeByID(id)
	if err != nil {
		return false, err
	}
	employeeRemoved := db.DB.Unscoped().Delete(&employee)
	if employeeRemoved.Error != nil {
		return false, employeeRemoved.Error
	}
	return true, nil
}
