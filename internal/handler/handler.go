package handler

import (
	"net/http"
	"sentraponsel-product-generator/internal/presenter"
	"sentraponsel-product-generator/service"

	"github.com/gofiber/fiber/v2"
)

func Generate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := []presenter.Request{}
		err := c.BodyParser(&body)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"Code": http.StatusBadRequest,
				"Error": err.Error(),
				"Success": false,
			})
		}

		response := service.GenerateData(body)

		return c.JSON(response)
	}
}