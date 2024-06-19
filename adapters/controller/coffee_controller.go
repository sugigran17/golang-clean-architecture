package controller

import "github.com/gofiber/fiber/v2"

type CoffeeController struct {
}

func (lc *CoffeeController) CoffeeBean(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}
