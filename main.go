package main

import (
	"level_5/config"
	"level_5/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()

	config.Connect()

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "john" && pass == "doe" {
				return true
			}
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).SendString("Unauthorized")
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	app.Get("/department", controller.GetDepartment)
	app.Post("/department", controller.AddDepartment)
	app.Get("/department/:id", controller.GetDepartmentById)
	app.Delete("/department/:id", controller.DeleteDepartmentById)
	app.Put("/department/:id", controller.UpdateDepartment)
	app.Get("/employee", controller.GetEmployee)
	app.Post("/employee", controller.AddEmployee)
	app.Get("/employee/:id", controller.GetEmployeeById)
	app.Delete("/employee/:id", controller.DeleteEmployeeById)
	app.Put("/employee/:id", controller.UpdateEmployee)

	app.Listen(": 3000")
}
