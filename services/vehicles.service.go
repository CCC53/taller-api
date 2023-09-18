package services

import (
	"strconv"
	"taller-api/db"
	"taller-api/models"

	"gorm.io/gorm"
)

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

func ListVehicles(pageSize int, page int) ([]models.Vehicle, int64) {
	vehicles := []models.Vehicle{}
	var totalCount int64
	db.DB.Scopes(Paginate(pageSize, page)).Find(&vehicles)
	db.DB.Model(&vehicles).Count(&totalCount)
	return vehicles, totalCount
}

func ListVehiclesSelect() []models.SelectResponse {
	vehicles := []models.Vehicle{}
	db.DB.Find(&vehicles)
	return mapVehiclesSelect(vehicles)
}

func GetVehicleByID(id string) (models.Vehicle, error) {
	vehicle := models.Vehicle{}
	vehicleDB := db.DB.Preload("Service").First(&vehicle, "id = ?", id)
	if vehicleDB.Error != nil {
		return models.Vehicle{}, vehicleDB.Error
	}
	return vehicle, nil
}

func CreateVehicle(formData models.Vehicle) (models.Vehicle, error) {
	vehicle := formData
	var vehicleDB *gorm.DB = db.DB.Create(&vehicle)
	if vehicleDB.Error != nil {
		return models.Vehicle{}, vehicleDB.Error
	}
	return vehicle, nil
}

func UpdateVehicle(id string, formData models.Vehicle) (models.Vehicle, error) {
	vehicleFound, err := GetVehicleByID(id)
	if err != nil {
		return models.Vehicle{}, err
	}
	vehicleFound.Brand = formData.Brand
	vehicleFound.Chassis = formData.Chassis
	vehicleFound.Model = formData.Model
	vehicleFound.Type = formData.Type
	vehicleFound.Year = formData.Year
	vehicleFound.Motor = formData.Motor
	vehicleFound.Plate = formData.Plate
	vehicleFound.Owner = formData.Owner
	vehicleFound.EmailOwner = formData.EmailOwner
	db.DB.Save(&vehicleFound)
	return vehicleFound, nil
}

func DeleteVehicle(id string) (bool, error) {
	vehicle, err := GetVehicleByID(id)
	if err != nil {
		return false, err
	}
	vehicleDeleted := db.DB.Unscoped().Delete(&vehicle)
	if vehicleDeleted.Error != nil {
		return false, vehicleDeleted.Error
	}
	return true, nil
}
