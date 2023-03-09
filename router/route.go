package router

import (
	"sentraponsel-product-generator/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func Route(app fiber.Router) {
	app.Post("/sentraponsel/generate", handler.Generate())
	app.Post("/parastar/generate/sales", handler.GenerateSales())
	app.Post("/parastar/generate/store", handler.GenerateStore())
	app.Post("/parastar/update/store", handler.UpdateStore())
	app.Post("/parastar/generate/user", handler.GenerateUser())
	app.Post("/parastar/generate/activity", handler.GenerateActivity())
	// app.Static("/", "./app")
	app.Static("/", "./swaggerui")
}