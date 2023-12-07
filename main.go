package main

import (
	"app/config"
	"app/db"
	"app/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	db.SetupDB()
	db.SetupRedis()
	app := fiber.New()
	router.RegisterRouter(app)
	if err := app.Listen(config.Port); err != nil {
		log.Fatal(err)
	}
}
