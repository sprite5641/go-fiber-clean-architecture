package route

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sprite5641/go-fiber-clean-architecture/api/controller"
	"github.com/sprite5641/go-fiber-clean-architecture/bootstrap"
	"github.com/sprite5641/go-fiber-clean-architecture/domain"
	"github.com/sprite5641/go-fiber-clean-architecture/mongo"
	"github.com/sprite5641/go-fiber-clean-architecture/repository"
	"github.com/sprite5641/go-fiber-clean-architecture/usecase"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, c fiber.Router) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
		Env:            env,
	}
	v1 := c.Group("/v1")

	v1.Get("/profile", pc.Fetch)

}
