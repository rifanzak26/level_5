package controller

import (
	"level_5/config"
	"level_5/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetEmployee(c *fiber.Ctx) error {
	user := c.Locals("admin").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if name != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)

	}
	var employees []model.Employee

	config.Database.Find(&employees)
	return c.Status(200).JSON(employees)
}

func GetEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee model.Employee

	result := config.Database.Find(&employee, id)

	user := c.Locals("admin").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if name != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(map[string]string{
			"message": "Employee Not Found",
		})
	}

	return c.Status(200).JSON(employee)
}

func AddEmployee(c *fiber.Ctx) error {
	employee := new(model.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	user := c.Locals("admin").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if name != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	config.Database.Create(&employee)
	return c.Status(201).JSON(employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	employee := new(model.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	user := c.Locals("admin").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if name != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	config.Database.Where("id = ?", id).Updates(&employee)
	return c.Status(200).JSON(employee)
}

func DeleteEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee model.Employee

	result := config.Database.Delete(&employee, id)

	user := c.Locals("admin").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if name != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(map[string]string{
			"message": "Data Employee not found, please check again",
		})
	}

	return c.Status(200).JSON(map[string]string{
		"message": "Employee success deleted",
	})
}
