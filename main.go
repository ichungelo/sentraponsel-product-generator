package main

import (
	"fmt"
	"log"
	"os"
	"sentraponsel-product-generator/db"
	"sentraponsel-product-generator/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	db.ConnOpensearch()
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func main() {
	appHost, appPort := os.Getenv("API_HOST"), os.Getenv("API_PORT")
	app := fiber.New()

	app.Use(cors.New())

	router.Route(app)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", appHost, appPort)))
}
