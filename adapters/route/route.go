package route

import "github.com/gofiber/fiber/v2"

func SetUpRoute(app *fiber.App) {
	NewCoffeeRouter(app)
}
