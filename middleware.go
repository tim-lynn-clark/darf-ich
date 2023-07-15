package darfich

import (
	"github.com/gofiber/fiber/v2"
)

// Config represents configuration values for the darfich package.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	// Optional. Default: nil
	Next       func(c *fiber.Ctx) bool
	ContextKey string
}

// New Create a new middleware handler
func New(config Config) func(*fiber.Ctx) error {
	// Return new Fiber handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if config.Next != nil && config.Next(c) {
			return c.Next()
		}

		// TODO: Strip URL params
		// TODO: Strip UUIDs route

		// TODO: Pull current user out of the context
		currentUser := c.Locals(config.ContextKey)
		// TODO: Pull user roll from user

		method := c.Method()
		route := c.Path()

		// TODO: Generate keys for role+method+route

		// TODO: Search through rules using key for matching rule
		// TODO: If no rule is found, return 403 Forbidden

		// TODO: If rule is found, and action is Allow, return c.Next()
		// TODO: If rule is found, and action is Deny, return 403 Forbidden

		// Send 204 No Content
		return c.SendStatus(fiber.StatusNoContent)
	}
}
