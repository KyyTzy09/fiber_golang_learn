package controllers

import (
	"fiber/api/services"
	"fiber/common/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	data := services.GetUserService()
	return c.JSON(fiber.Map{
		"data": data,
		"status": fiber.Map{
			"message":     "getting users successfully",
			"status_code": fiber.StatusOK,
		},
	})
}

func GetUserById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	userId, err := strconv.Atoi(paramId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Id must be a number"})
	}
	user := services.GetUserByIdService(userId)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
		"status": fiber.Map{
			"message":     "getting user successfully",
			"status_code": fiber.StatusOK,
		},
	})
}

func CreateUser(c *fiber.Ctx) error {
	var body types.CreateUser

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	createdUser, err := services.CreateUserService(body)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data": createdUser,
		"status": fiber.Map{
			"message":     "user created successfully",
			"status_code": fiber.StatusCreated,
		},
	})
}
