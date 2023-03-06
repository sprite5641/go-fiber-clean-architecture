package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sprite5641/go-fiber-clean-architecture/bootstrap"
	"github.com/sprite5641/go-fiber-clean-architecture/domain"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domain.RefreshTokenUsecase
	Env                 *bootstrap.Env
}

func (rtc *RefreshTokenController) RefreshToken(c *fiber.Ctx) error {
	var request domain.RefreshTokenRequest
	ctx := c.Context()
	err := c.BodyParser(&request)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	id, err := rtc.RefreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByID(ctx, id)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	accessToken, err := rtc.RefreshTokenUsecase.CreateAccessToken(&user, rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	refreshToken, err := rtc.RefreshTokenUsecase.CreateRefreshToken(&user, rtc.Env.RefreshTokenSecret, rtc.Env.RefreshTokenExpiryHour)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return c.JSON(fiber.Map{"error": refreshTokenResponse})
}
