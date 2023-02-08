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

func GenerateSales() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var(
			body = presenter.RequestSales{}
			err = c.BodyParser(&body)
		)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"Code": http.StatusBadRequest,
				"Error": err.Error(),
				"Success": false,
			})
		}

		response := service.GenerateDataSales(body)

		return c.JSON(response)
	}
}

func GenerateStore() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			body = presenter.RequestStore{}
			err = c.BodyParser(&body)
		)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"Code": http.StatusBadRequest,
				"Error": err.Error(),
				"Success": false,
			})
		}

		response := service.GenerateDataStore(body)

		return c.JSON(response)
	}
}

func UpdateStore() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			body = presenter.RequestStore{}
			err = c.BodyParser(&body)
		)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"Code": http.StatusBadRequest,
				"Error": err.Error(),
				"Success": false,
			})
		}

		response := service.UpdateDataStore(body)

		return c.JSON(response)
	}
}

func GenerateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			body = presenter.RequestUser{}
			err = c.BodyParser(&body)
		)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"Code": http.StatusBadRequest,
				"Error": err.Error(),
				"Success": false,
			})
		}

		response := service.GenerateUser(body)

		return c.JSON(response)
	}
}