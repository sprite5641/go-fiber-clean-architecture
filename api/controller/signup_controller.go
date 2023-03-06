package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sprite5641/go-fiber-clean-architecture/bootstrap"
	"github.com/sprite5641/go-fiber-clean-architecture/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c *fiber.Ctx) error {
	var request domain.SignupRequest
	ctx := c.Context()
	err := c.BodyParser(&request)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})

	}

	_, err = sc.SignupUsecase.GetUserByEmail(ctx, request.Email)
	if err == nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})

	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(ctx, &user)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		return c.JSON(fiber.Map{"error": domain.ErrorResponse{Message: err.Error()}})
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.JSON(fiber.Map{"error": signupResponse})

}
