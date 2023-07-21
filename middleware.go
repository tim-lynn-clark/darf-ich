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

		// Search through rules using key for matching rule
		var valid bool
		for _, rule := range config.RuleSet.Rules {
			if fiber.RoutePatternMatch(path, string(rule.Route)) &&
				rule.Method == utils.HttpMethod(method) &&
				rule.Role == utils.Role(currentUser.RoleName) {

				if rule.Action == utils.ActionAllow {
					valid = true
				}
			}
		}

		if valid {
			return c.Next()
		}
		// If no rule is found, return 403 Forbidden
		return c.SendStatus(fiber.StatusForbidden)
	}
}
