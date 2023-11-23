package routes

import (
	"taller-api/controllers"
	"taller-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func authRouter(router fiber.Router) {
	router.Post("/login", controllers.Login)
	router.Get("/validate-token", middlewares.IsAuthenticated, controllers.ValidateToken)
	router.Get("/validate-role", middlewares.IsAuthenticated, controllers.ValidateRole)
	router.Get("/me", middlewares.IsAuthenticated, controllers.GetMe)
	router.Put("/update-profile", middlewares.IsAuthenticated, controllers.UpdateProfile)
	router.Get("/logout", controllers.Logout)
}

func employeesRouter(router fiber.Router) {
	router.Get("/", middlewares.IsAdmin, controllers.FindEmployees)
	router.Get("/:id", middlewares.IsAdmin, controllers.FindEmployee)
	router.Post("/", middlewares.IsAdmin, controllers.CreateEmployee)
	router.Put("/:id", middlewares.IsAdmin, controllers.UpdateEmployee)
	router.Delete("/:id", middlewares.IsAdmin, controllers.DeleteEmployee)
}

func vehiclesRouter(router fiber.Router) {
	router.Get("/", controllers.FindVehicles)
	router.Get("/select", controllers.FindVehiclesSelect)
	router.Get("/:id", controllers.FindVehicle)
	router.Post("/", middlewares.IsAdmin, controllers.CreateVehicle)
	router.Put("/:id", middlewares.IsAdmin, controllers.UpdateVehicle)
	router.Delete("/:id", middlewares.IsAdmin, controllers.DeleteVehicle)
}

func sparePartsRouter(router fiber.Router) {
	router.Get("/", controllers.FindSpareParts)
	router.Get("/:id", controllers.FindSparePart)
	router.Post("/", middlewares.IsAdmin, controllers.CreateSparePart)
	router.Put("/:id", middlewares.IsAdmin, controllers.UpdateSparePart)
	router.Delete("/:id", middlewares.IsAdmin, controllers.DeleteSparePart)
}

func servicesRouter(router fiber.Router) {
	router.Get("/", controllers.FindServices)
	router.Get("/:id", controllers.FindService)
	router.Post("/", controllers.CreateService)
	router.Put("/:id", controllers.UpdateService)
	router.Delete("/:id", controllers.DeleteService)
}

func assignmentsRouter(router fiber.Router) {
	router.Get("/employees", controllers.FindMechanicsAviable)
	router.Get("/spare-parts", controllers.FindSaprePartsAviable)
	router.Put("/employee/:id", controllers.AssignEmployeeToService)
	router.Put("/spare-part/:id", controllers.AssignSparePartToService)
	router.Delete("/employee/:id", controllers.RemoveEmployeeFromService)
	router.Delete("/spare-part/:id", controllers.RemoveSparePartFromService)
}

func searchesRouter(router fiber.Router) {
	router.Get("/:table", middlewares.IsAuthenticated, controllers.SearchByTable)
	router.Get("/restore/:table", middlewares.IsAuthenticated, controllers.ResotreData)
}

func InitRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	authRouter(auth)

	employees := api.Group("/employees", middlewares.IsAuthenticated)
	employeesRouter(employees)

	vehicles := api.Group("/vehicles", middlewares.IsAuthenticated)
	vehiclesRouter(vehicles)

	spareParts := api.Group("/spare-parts", middlewares.IsAuthenticated)
	sparePartsRouter(spareParts)

	services := api.Group("/services", middlewares.IsAuthenticated)
	servicesRouter(services)

	assign := api.Group("/assignments", middlewares.IsAuthenticated)
	assignmentsRouter(assign)

	searches := api.Group("/search")
	searchesRouter(searches)
}
