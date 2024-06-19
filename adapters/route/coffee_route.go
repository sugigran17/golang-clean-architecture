package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugigran17/golang-clean-architecture/adapters/controller"
)

func NewCoffeeRouter(app *fiber.App) {
	api := app.Group("/coffee")
	lc := &controller.CoffeeController{}

	api.Get("/bean", lc.CoffeeBean)
}
