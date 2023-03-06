package controller

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"github.com/sprite5641/go-fiber-clean-architecture/bootstrap"
	"github.com/sprite5641/go-fiber-clean-architecture/domain"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c *fiber.Ctx) error {
	var request domain.LoginRequest
	ctx := c.Context()
	err := c.BodyParser(&request)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	user, err := lc.LoginUsecase.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.JSON(fiber.Map{"error": loginResponse})

}
