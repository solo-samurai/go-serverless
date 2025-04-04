package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/solo-samurai/go-serverless/app/routes"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env")
	}

	app := fiber.New()
	routes.Register(app)

	app.Listen(":8080")
}
