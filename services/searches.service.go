package services

import (
	"taller-api/db"
	"taller-api/models"
)

func Search(table string, search string, page int, pageSize int) map[string]interface{} {
	response := make(map[string]interface{})
	regex := "%" + search + "%"
	var totalCount int64
	switch table {
	case "employees":
		employees := []models.Employee{}
		db.DB.Scopes(Paginate(pageSize, page)).Where("name ILIKE ?", regex).Find(&employees)
		db.DB.Model(&employees).Where("name ILIKE ?", regex).Count(&totalCount)
		response["data"] = MappingEmployees(employees)
		response["total"] = totalCount
	case "services":
		services := []models.Service{}
		db.DB.Scopes(Paginate(pageSize, page)).Where("issue ILIKE ?", regex).Find(&services)
		db.DB.Model(&services).Where("issue ILIKE ?", regex).Count(&totalCount)
		response["data"] = services
		response["total"] = totalCount
	case "spare-parts":
		spareParts := []models.SparePart{}
		db.DB.Scopes(Paginate(pageSize, page)).Where("name ILIKE ? OR supplier ILIKE ?", regex, regex).Find(&spareParts)
		db.DB.Model(&spareParts).Where("name ILIKE ? OR supplier ILIKE ?", regex, regex).Count(&totalCount)
		response["data"] = spareParts
		response["total"] = totalCount
	case "vehicles":
		vehicles := []models.Vehicle{}
		db.DB.Scopes(Paginate(pageSize, page)).Where("model ILIKE ? OR brand ILIKE ?", regex, regex).Find(&vehicles)
		db.DB.Model(&vehicles).Where("model ILIKE ? OR brand ILIKE ?", regex, regex).Count(&totalCount)
		response["data"] = vehicles
		response["total"] = totalCount
	}
	return response
}

func ResotreData(table string, pageSize int, page int, token string) map[string]interface{} {
	response := make(map[string]interface{})
	switch table {
	case "employees":
		data, totalCount := ListEmployees(pageSize, page, token)
		response["data"] = data
		response["total"] = totalCount
	case "vehicles":
		data, totalCount := ListVehicles(pageSize, page)
		response["data"] = data
		response["total"] = totalCount
	case "services":
		data, totalCount := ListServices(pageSize, page)
		response["data"] = data
		response["total"] = totalCount
	case "spare-parts":
		data, totalCount := ListSpareParts(pageSize, page)
		response["data"] = data
		response["total"] = totalCount
	}
	return response
}
