package main

import (
	"fmt"
	"log"
	"sentraponsel-product-generator/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appHost, appPort := "localhost", "5000" 
	app := fiber.New()

	router.Route(app)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", appHost, appPort)))
}