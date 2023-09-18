package services

import (
	"taller-api/db"
	"taller-api/models"

	"gorm.io/gorm"
)

func ListSpareParts(pageSize int, page int) ([]models.SparePart, int64) {
	spareParts := []models.SparePart{}
	var totalCount int64
	db.DB.Scopes(Paginate(pageSize, page)).Find(&spareParts)
	db.DB.Model(&spareParts).Count(&totalCount)
	return spareParts, totalCount
}

func GetSparePartByID(id string) (models.SparePart, error) {
	sparePart := models.SparePart{}
	sparePartDB := db.DB.First(&sparePart, "id = ?", id)
	if sparePartDB.Error != nil {
		return models.SparePart{}, sparePartDB.Error
	}
	return sparePart, nil
}

func CreateSparePart(formData models.SparePart) (models.SparePart, error) {
	sparePart := formData
	var sparePartDB *gorm.DB = db.DB.Create(&sparePart)
	if sparePartDB.Error != nil {
		return models.SparePart{}, sparePartDB.Error
	}
	return sparePart, nil
}

func UpdateSparePart(id string, formData models.SparePart) (models.SparePart, error) {
	sparePartFound, err := GetSparePartByID(id)
	if err != nil {
		return models.SparePart{}, err
	}
	sparePartFound.Name = formData.Name
	sparePartFound.Disponible = formData.Disponible
	sparePartFound.Price = formData.Price
	sparePartFound.Supplier = formData.Supplier
	sparePartFound.PurchaseDate = formData.PurchaseDate
	sparePartFound.Type = formData.Type
	db.DB.Save(&sparePartFound)
	return sparePartFound, nil
}

func DeleteSparePart(id string) (bool, error) {
	sparePart, err := GetSparePartByID(id)
	if err != nil {
		return false, err
	}
	sparePartDeleted := db.DB.Unscoped().Delete(&sparePart)
	if sparePartDeleted.Error != nil {
		return false, sparePartDeleted.Error
	}
	return true, nil
}
