package main

import (
	"level_5/config"
	"level_5/controller"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	app := fiber.New()

	config.Connect()

	app.Post("/login", controller.Login)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"), ContextKey: "admin",
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

func Restricted(c *fiber.Ctx) error {
	user := c.Locals("admin").(jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
