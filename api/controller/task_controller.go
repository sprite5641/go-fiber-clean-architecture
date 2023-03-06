package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sprite5641/go-fiber-clean-architecture/bootstrap"
	"github.com/sprite5641/go-fiber-clean-architecture/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
	Env         *bootstrap.Env
}

func (tc *TaskController) Create(c *fiber.Ctx) error {
	var task domain.Task
	ctx := c.Context()

	err := c.BodyParser(&task)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})

	}

	userID := c.Locals("userID").(string)
	task.ID = primitive.NewObjectID()

	task.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	err = tc.TaskUsecase.Create(ctx, &task)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
}

func (u *TaskController) Fetch(c *fiber.Ctx) error {
	userID := c.Locals("x-user-id").(string)
	ctx := c.Context()

	tasks, err := u.TaskUsecase.FetchByUserID(ctx, userID)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})

	}
	return c.JSON(fiber.Map{"error": tasks})
}
