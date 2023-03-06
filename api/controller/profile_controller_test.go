package controller_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/sprite5641/go-fiber-clean-architecture/api/controller"
	"github.com/sprite5641/go-fiber-clean-architecture/domain"
	"github.com/sprite5641/go-fiber-clean-architecture/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setUserID(userID string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("userID", userID)
		return c.Next()
	}
}

func TestFetch(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockProfile := &domain.Profile{
			Name:  "Test Name",
			Email: "test@gmail.com",
		}

		userObjectID := primitive.NewObjectID()
		userID := userObjectID.Hex()

		mockProfileUsecase := new(mocks.ProfileUsecase)

		mockProfileUsecase.On("GetProfileByID", mock.Anything, userID).Return(mockProfile, nil)

		app := fiber.New()

		pc := &controller.ProfileController{
			ProfileUsecase: mockProfileUsecase,
		}

		app.Use(setUserID(userID))
		app.Get("/profile", pc.Fetch)

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var body domain.Profile
		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			t.Fatal(err)
		}

		err = json.Unmarshal(bodyBytes, &body)
		assert.NoError(t, err)

		assert.Equal(t, mockProfile.Name, body.Name)
		assert.Equal(t, mockProfile.Email, body.Email)

		mockProfileUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		userObjectID := primitive.NewObjectID()
		userID := userObjectID.Hex()

		mockProfileUsecase := new(mocks.ProfileUsecase)

		customErr := errors.New("Unexpected")

		mockProfileUsecase.On("GetProfileByID", mock.Anything, userID).Return(nil, customErr)

		app := fiber.New()

		pc := &controller.ProfileController{
			ProfileUsecase: mockProfileUsecase,
		}

		app.Use(setUserID(userID))
		app.Get("/profile", pc.Fetch)

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		var body domain.ErrorResponse
		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			t.Fatal(err)
		}
		err = json.Unmarshal(bodyBytes, &body)
		assert.NoError(t, err)

		assert.Equal(t, customErr.Error(), body.Message)

		mockProfileUsecase.AssertExpectations(t)
	})

}
