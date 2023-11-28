package main

import (
	"app/config"
	"app/core/database"
	"app/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	database.SetupDB()
	database.SetupRedis()
	app := fiber.New()
	router.RegisterRouter(app)
	if err := app.Listen(config.Port); err != nil {
		log.Fatal(err)
	}
}
