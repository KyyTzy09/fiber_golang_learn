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
	user, err := services.GetUserByIdService(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

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
			"status": fiber.StatusInternalServerError,
			"error":  err.Error(),
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

func UpdateUser(c *fiber.Ctx) error {
	var body types.UpdateUserRequest
	var params = c.Params("id")

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	createdUser, err := services.UpdateUser(params, body.UserName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  err.Error(),
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

func DeleteUserById(c *fiber.Ctx) error {

	deletedUsers, err := services.DeleteUserById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": deletedUsers,
		"status": fiber.Map{
			"message":     "user deleted successfully",
			"status_code": 200,
		},
	})
}

func DeleteAllUsers(c *fiber.Ctx) error {

	deletedUsers, err := services.DeleteAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": deletedUsers,
		"status": fiber.Map{
			"message":     "user deleted successfully",
			"status_code": fiber.StatusAccepted,
		},
	})
}
