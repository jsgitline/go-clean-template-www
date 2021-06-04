package main

import (
	"clean-template/settings"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	run()
}

func run() {

	// Инициализация приложения
	// init fiber app
	app := settings.SetupFiber()

	// Чтение потров
	// Listen ports
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
