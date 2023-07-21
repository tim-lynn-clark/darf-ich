package darfich

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tim-lynn-clark/darfich/ability"
	"github.com/tim-lynn-clark/darfich/utils"
)

// Config represents configuration values for the DarfIch package.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	// Optional. Default: nil
	Next       func(c *fiber.Ctx) bool
	ContextKey string
	RuleSet    ability.Set
}

// New Create a new middleware handler
func New(config Config) func(*fiber.Ctx) error {
	// Return new Fiber handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if config.Next != nil && config.Next(c) {
			return c.Next()
		}

		// Pull current user out of the context
		currentUser := c.Locals(config.ContextKey).(utils.DtoCurrentUser)

		method := c.Method()
		path := c.Path()

		// Generate keys for role+method+route+resource
		_, hashKey := ability.GenerateRuleKeys(
			utils.Role(currentUser.RoleName),
			utils.HttpMethod(method),
			utils.HttpRoute(path))

		// Search through rules using key for matching rule
		for _, rule := range config.RuleSet.Rules {
			if rule.HashKey == hashKey {
				if rule.Action == utils.ActionAllow {
					return c.Next()
				} else {
					return c.SendStatus(fiber.StatusForbidden)
				}
			}
		}

		// If no rule is found, return 403 Forbidden
		return c.SendStatus(fiber.StatusForbidden)
	}
}
