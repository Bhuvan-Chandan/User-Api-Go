package handler

import (
	"go-user-api/internal/logger"
	"go-user-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	Service service.UserServiceInterface
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	user, err := h.Service.GetUserWithAge(int32(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "user not found",
		})
	}

	return c.JSON(user)
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob" validate:"required"`
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest

	// 1. Parse request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "invalid request",
		})
	}

	// 2. Validate request
	if err := validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "validation failed",
			"error":   err.Error(),
		})
	}

	// 3. Call service (DB insert)
	user, err := h.Service.CreateUser(req.Name, req.DOB)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to create user",
			"error":   err.Error(),
		})
	}

	// 4. Log (DO NOT return this)
	logger.Log.Info("user created",
		zap.String("name", req.Name),
	)

	// 5. Return response
	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "user created",
		"data":    user,
	})
}

type UpdateUserRequest struct {
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	users, err := h.Service.ListUsers(page, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to fetch users",
			"error":   err.Error(),
		})
	}

	logger.Log.Info("users fetched")

	return c.JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {

	id, _ := c.ParamsInt("id")

	var req UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "invalid request",
		})
	}

	user, err := h.Service.UpdateUser(int32(id), req.Name, req.DOB)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to update user",
			"error":   err.Error(),
		})
	}

	logger.Log.Info("user updated",
		zap.Int("id", id),
	)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "user updated",
		"data":    user,
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {

	id, _ := c.ParamsInt("id")

	err := h.Service.DeleteUser(int32(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to delete user",
			"error":   err.Error(),
		})
	}

	logger.Log.Info("user deleted",
		zap.Int("id", id),
	)

	return c.SendStatus(204)
}

var validate = validator.New()
