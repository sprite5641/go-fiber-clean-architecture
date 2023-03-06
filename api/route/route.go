package route

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sprite5641/go-fiber-clean-architecture/api/middleware"
	"github.com/sprite5641/go-fiber-clean-architecture/bootstrap"
	"github.com/sprite5641/go-fiber-clean-architecture/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, c *fiber.App) {

	// hello world
	c.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	publicRouter := c.Group("/api")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := c.Group("/api")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)

}
