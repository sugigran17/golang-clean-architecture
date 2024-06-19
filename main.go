package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sugigran17/golang-clean-architecture/adapters/route"
	"github.com/sugigran17/golang-clean-architecture/database"
	"github.com/sugigran17/golang-clean-architecture/repository"
	"os"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "golang-clean-architecture",
	})

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", DB_USER, DB_PASS, DB_HOST, DB_PORT)
	mongoClient := database.NewMongoClient(mongodbURI, 10)
	repository.NewCoffeeRepository(mongoClient)

	route.SetUpRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Golang Clean-Architecture!")
	})

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	app.Listen(":" + appPort)
}
