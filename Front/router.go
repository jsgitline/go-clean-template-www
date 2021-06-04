package Front

import (
	"clean-template/Front/handlers"
	"github.com/gofiber/fiber/v2"

)

func Router(app *fiber.App) {
	app.Static("/public/css/", "./Front/public/css")
	app.Static("/public/js/", "./Front/public/js")
	app.Static("/public/img/", "./Front/public/img")

	app.Get("/", handlers.FrontPage)
}
