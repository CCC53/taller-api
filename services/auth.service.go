package services

import (
	"errors"
	"taller-api/config"
	"taller-api/db"
	"taller-api/enums"
	"taller-api/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func FindByCredentials(email string, password string) (*models.Employee, error) {
	var employee *models.Employee
	employeeDB := db.DB.First(&employee, "email = ?", email)
	if employeeDB.Error != nil {
		return nil, errors.New("email or password incorrect")
	}
	if !ComparePassword(employee.Password, password) {
		return nil, errors.New("email or password incorrect")
	}
	return employee, nil
}

func GenerateToken(employee models.Employee) (string, error) {
	claims := jwt.MapClaims{
		"id":    employee.ID,
		"name":  employee.Name,
		"email": employee.Email,
		"role":  employee.Role,
		"exp":   time.Now().Add((time.Hour) * 1).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(config.Secret))
	return token, err
}

func LoadMenu(role enums.ValidRoles) []models.Routes {
	routes := []models.Routes{
		{Route: "employees", Icon: "groups", Label: "Empleados"},
		{Route: "vehicles", Icon: "directions_car", Label: "Vehículos"},
		{Route: "spare-parts", Icon: "settings", Label: "Refacciones"},
		{Route: "services", Icon: "car_crash", Label: "Servicios"},
		{Route: "update-profile", Icon: "account_circle", Label: "Perfil"},
	}
	if role == enums.Mechanic {
		routes = routes[1:]
	}
	return routes
}

func ValidateToken(token string) bool {
	var current = time.Now().Unix()
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		return false
	}
	return float64(current) < claims["exp"].(float64)
}

func ValidateRole(token string) bool {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		return false
	}
	return claims["role"] == string(enums.Admin)
}

func GetMe(token string) *models.Employee {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		return nil
	}
	employee, _ := GetEmployeeByID(claims["id"].(string))
	return &employee
}

func GetMyID(token string) string {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		return ""
	}
	return claims["id"].(string)
}

func UpdateProfile(token string, formData models.Employee) models.EmployeeResponse {
	id := GetMyID(token)
	updated, _ := UpdateEmployee(id, formData)
	return updated
}
