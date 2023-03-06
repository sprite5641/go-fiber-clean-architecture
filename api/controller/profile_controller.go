package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sprite5641/go-fiber-clean-architecture/bootstrap"
	"github.com/sprite5641/go-fiber-clean-architecture/domain"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
	Env            *bootstrap.Env
}

func (pc *ProfileController) Fetch(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	ctx := c.Context()
	profile, err := pc.ProfileUsecase.GetProfileByID(ctx, userID)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})

	}

	return c.JSON(fiber.Map{"error": profile})

}
