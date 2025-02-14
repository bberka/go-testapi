package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// ParseBody is a generic request parser middleware
func ParseBody[T any](handler func(c *fiber.Ctx, req *T) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req T

		// Parse JSON body
		if err := ctx.BodyParser(&req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Pass parsed request to the next handler
		return handler(ctx, &req)
	}
}

// ParseBodyWithValidation (Auto Parses & Validates Request)
func ParseBodyWithValidation[T any](handler func(c *fiber.Ctx, req *T) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req T

		// Parse JSON body
		if err := ctx.BodyParser(&req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Validate struct if it has a Validate method
		if validator, ok := any(&req).(interface{ Validate() error }); ok {
			if err := validator.Validate(); err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			}
		}

		// Pass parsed request to the next handler
		return handler(ctx, &req)
	}
}
