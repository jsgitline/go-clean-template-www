package settings

import (
	"clean-template/Front"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

// Инициализируем роуты приложений
// Init apps routes
func SetupFiber() *fiber.App {

	// run new fiber object
	app := fiber.New(fiber.Config{
		Views: html.New("./templates", ".html"),
	})

	// Register the index route with a simple
	// "OK" response. It should return status
	// code 200

	// ROUTES FRONT
	Front.Router(app)

	return app
}
