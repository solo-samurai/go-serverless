package middleware

import (
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gofiber/fiber/v2"
)

func RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := clerk.SessionClaimsFromContext(c.Context())
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		c.Locals("user", claims)

		return c.Next()
	}

}
