package controller

import (
	"level_5/config"
	"level_5/model"

	"github.com/gofiber/fiber/v2"
)

func GetDepartment(c *fiber.Ctx) error {
	var departments []model.Department

	config.Database.Find(&departments)
	return c.Status(200).JSON(departments)
}

func GetDepartmentById(c *fiber.Ctx) error {
	id := c.Params("id")
	var department model.Department

	result := config.Database.Find(&department, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(map[string]string{
			"message": "Department Not Found",
		})
	}

	return c.Status(200).JSON(department)
}

func AddDepartment(c *fiber.Ctx) error {
	department := new(model.Department)

	if err := c.BodyParser(department); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&department)
	return c.Status(201).JSON(department)
}

func UpdateDepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	department := new(model.Department)

	if err := c.BodyParser(department); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&department)
	return c.Status(200).JSON(department)
}

func DeleteDepartmentById(c *fiber.Ctx) error {
	id := c.Params("id")

	var department model.Department

	result := config.Database.Delete(&department, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(map[string]string{
			"message": "Data Department not found, please check again",
		})
	}

	return c.Status(200).JSON(map[string]string{
		"message": "Department success deleted",
	})
}
