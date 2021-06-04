package handlers

import "github.com/gofiber/fiber/v2"

// Example handler
func FrontPage(c *fiber.Ctx) error {
	return c.Render("base", fiber.Map{})
}
