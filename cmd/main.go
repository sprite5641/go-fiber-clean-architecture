package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sprite5641/go-fiber-clean-architecture/api/route"
	"github.com/sprite5641/go-fiber-clean-architecture/bootstrap"
)

func main() {

	appInit := bootstrap.App()

	env := appInit.Env

	db := appInit.Mongo.Database(env.DBName)
	defer appInit.CloseDBConnection()
	// KDFYDzgwq4ZKHRFh mongo-test1

	timeout := time.Duration(env.ContextTimeout) * time.Second

	app := fiber.New()

	route.Setup(env, timeout, db, app)

	app.Listen(":8080")
}
