package services

import (
	"taller-api/db"
	"taller-api/models"

	"gorm.io/gorm"
)

func mapServicePopulated(service models.Service) models.ServiceResponse {
	serviceMapped := models.ServiceResponse{
		ID:         service.ID,
		StartDate:  service.StartDate,
		EndDate:    service.EndDate,
		Issue:      service.Issue,
		VehicleID:  *service.VehicleID,
		Employees:  MappingEmployees(*service.Employees),
		SpareParts: *service.SpareParts,
	}
	return serviceMapped
}

func getServiceByID(id string) (models.Service, error) {
	service := models.Service{}
	serviceDB := db.DB.First(&service, "id = ?", id)
	if serviceDB.Error != nil {
		return models.Service{}, serviceDB.Error
	}
	return service, nil
}

func ListServices(pageSize int, page int) ([]models.Service, int64) {
	services := []models.Service{}
	var totalCount int64
	db.DB.Scopes(Paginate(pageSize, page)).Find(&services)
	db.DB.Model(&services).Count(&totalCount)
	return services, totalCount
}

func GetServicePopulatedByID(id string) (models.ServiceResponse, error) {
	service := models.Service{}
	serviceDB := db.DB.First(&service, "id = ?", id)
	if serviceDB.Error != nil {
		return models.ServiceResponse{}, serviceDB.Error
	}
	db.DB.Model(&service).Association("Employees").Find(&service.Employees)
	db.DB.Model(&service).Association("SpareParts").Find(&service.SpareParts)
	return mapServicePopulated(service), nil
}

func CreateService(formData models.Service) (models.Service, error) {
	service := formData
	var serviceDB *gorm.DB = db.DB.Create(&service)
	if serviceDB.Error != nil {
		return models.Service{}, serviceDB.Error
	}
	return service, nil
}

func UpdateService(id string, formData models.Service) (models.ServiceResponse, error) {
	serviceDB, err := getServiceByID(id)
	if err != nil {
		return models.ServiceResponse{}, err
	}
	serviceDB.StartDate = formData.StartDate
	serviceDB.EndDate = formData.EndDate
	serviceDB.Issue = formData.Issue
	serviceDB.VehicleID = formData.VehicleID
	db.DB.Save(&serviceDB)
	servicePopulated, _ := GetServicePopulatedByID(id)
	return servicePopulated, nil
}

func DeleteService(id string) (bool, error) {
	serviceDB, err := getServiceByID(id)
	if err != nil {
		return false, err
	}
	serviceDeleted := db.DB.Unscoped().Delete(&serviceDB)
	if serviceDeleted.Error != nil {
		return false, serviceDeleted.Error
	}
	return true, nil
}
