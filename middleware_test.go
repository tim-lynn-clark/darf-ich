package darfich

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/tim-lynn-clark/darfich/ability"
	"github.com/tim-lynn-clark/darfich/utils"
)

func TestNew(t *testing.T) {
	t.Parallel()

	const contextKey = "current_user"
	currentUser := utils.DtoCurrentUser{
		ID:       uuid.New(),
		Email:    "test@criticalprep.com",
		RoleID:   uuid.New(),
		RoleName: "admin",
	}

	ruleSet := ability.Set{
		Rules:       []ability.Rule{},
		Credentials: []utils.Credential{},
	}

	_, err := ruleSet.Can(
		utils.Role(currentUser.RoleName),
		utils.HttpGet,
		utils.HttpRoute("/book/:id"),
		utils.Resource("book"),
	)
	if err != nil {
		t.Errorf("ruleSet.Can() error = %v", err)
	}

	config := Config{
		Next:       nil,
		ContextKey: contextKey,
		RuleSet:    ruleSet,
	}

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals(contextKey, currentUser)
		return c.Next()
	})

	// Test DarfIch middleware
	app.Use(New(config))

	// 403 Forbidden
	app.Get("/test/:user", func(c *fiber.Ctx) error {
		utils.AssertEqual(t, "/test/John", c.Path())
		return nil
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/test/john", nil))
	utils.AssertEqual(t, fiber.StatusForbidden, resp.StatusCode, "Status Forbidden Code")

	// 200 Success
	app.Get("/book/:id", func(c *fiber.Ctx) error {
		return nil
	})

	resp, err = app.Test(httptest.NewRequest("GET", "/book/1243", nil))
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status OK (Success) Code")
}
