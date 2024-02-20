package controller

import (
	"level_5/config"
	"level_5/model"

	"github.com/gofiber/fiber/v2"
)

func GetEmployee(c *fiber.Ctx) error {
	var employees []model.Employee

	config.Database.Find(&employees)
	return c.Status(200).JSON(employees)
}

func GetEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee model.Employee

	result := config.Database.Find(&employee, id)

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

	config.Database.Create(&employee)
	return c.Status(201).JSON(employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	employee := new(model.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&employee)
	return c.Status(200).JSON(employee)
}

func DeleteEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee model.Employee

	result := config.Database.Delete(&employee, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(map[string]string{
			"message": "Data Employee not found, please check again",
		})
	}

	return c.Status(200).JSON(map[string]string{
		"message": "Employee success deleted",
	})
}
