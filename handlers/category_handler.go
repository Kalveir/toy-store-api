package handler

import (
	"api/models"
	"api/service"
	"strconv"
	"github.com/gofiber/fiber/v3"
)

/*
HANDLER GET CATEGORY
*/
func GetCategory(c fiber.Ctx) error {
	category := service.GetAllCategory()
	return c.Status(200).JSON(category)
}

/*
HANDLER FIND CATEGORY
*/
func FindCategory(c fiber.Ctx) error {
	id := c.Params("id")
	category := service.GetCategoryById(id)
	return c.Status(200).JSON(category)
}

/*
HANDLER CREATE CATEGORY
*/
func CreateCategory(c fiber.Ctx) error {
	var category models.Category
	if err := c.Bind().Body(&category); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	categories := service.CreateCategory(category)
	return c.Status(200).JSON(categories)
}

/*
HANDLER UPDATE CATEGORY
*/
func UpdateCategory(c fiber.Ctx) error {
	id,_ := strconv.Atoi(c.Params("id"))
	var category models.Category
	if err := c.Bind().Body(&category); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	service.UpdateCategory(uint(id), category)
	return c.Status(200).JSON(category)
}

/*
HANDLER DELETE CATEGORY
*/
func DeleteCategory(c fiber.Ctx) error {
	id,_ := strconv.Atoi(c.Params("id"))
	if err := service.DeleteCategory(uint(id)); err != nil {
		return c.Status(500).SendString("Failed Delete Category")
	}
	return c.Status(200).JSON(fiber.Map{
		"message":"Success",
	})
}
