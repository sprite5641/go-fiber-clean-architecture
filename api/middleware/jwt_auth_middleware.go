package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sprite5641/go-fiber-clean-architecture/domain"
	"github.com/sprite5641/go-fiber-clean-architecture/internal/tokenutil"
)

func JwtAuthMiddleware(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					return c.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{Message: err.Error()})
				}
				c.Set("x-user-id", userID)
				return c.Next()
			}
			return c.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{Message: err.Error()})
		}
		return c.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{Message: "Not authorized"})
	}
}
