package router

import (
	"sentraponsel-product-generator/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func Route(app fiber.Router) {
	app.Post("/generate", handler.Generate())
}