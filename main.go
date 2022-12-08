package main

import (
	"fmt"
	"log"
	"sentraponsel-product-generator/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	appHost, appPort := "localhost", "5001" 
	app := fiber.New()

	app.Use(cors.New())

	router.Route(app)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", appHost, appPort)))
}