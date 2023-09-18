package enums

type ValidRoles string

const (
	Admin    ValidRoles = "admin"
	Mechanic ValidRoles = "mechanic"
)

type ValidSpareParts string

const (
	Maintenance ValidSpareParts = "maintenance"
	Replacement ValidSpareParts = "replacement"
)

type ValidVehicleTypes string

const (
	Sedan   ValidVehicleTypes = "sedan"
	Torton  ValidVehicleTypes = "toron"
	Trailer ValidVehicleTypes = "trailer"
	SUV     ValidVehicleTypes = "suv"
	Pickup  ValidVehicleTypes = "pickup"
)
