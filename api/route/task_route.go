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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, c fiber.Router) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
		Env:         env,
	}

	v1 := c.Group("/v1")

	v1.Get("/tasks", tc.Fetch)
}
