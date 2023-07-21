package ability

import (
	"errors"
	"fmt"
	"testing"

	"github.com/tim-lynn-clark/darfich/utils"
)

func TestCan(t *testing.T) {
	testCases := []struct {
		name         string
		role         utils.Role
		method       utils.HttpMethod
		route        utils.HttpRoute
		resource     utils.Resource
		expStringKey string
		expHashKey   string
		expError     error
	}{
		{
			name:         "Rule should be generated successfully",
			role:         "admin",
			method:       "GET",
			route:        "/books",
			resource:     "book",
			expStringKey: "admin | GET | /books | book",
			expHashKey:   "682af11c5a0f6fe88c9799a215fb41128b34c89c811510ca62d6a8f5c7230592",
			expError:     nil,
		},
		{
			name:         "Duplicate rule should NOT be generated",
			role:         "admin",
			method:       "GET",
			route:        "/books",
			resource:     "book",
			expStringKey: "admin | GET | /books | book",
			expHashKey:   "682af11c5a0f6fe88c9799a215fb41128b34c89c811510ca62d6a8f5c7230592",
			expError:     &utils.ExistingRuleError{StringKey: "admin | GET | /books"},
		},
	}

	ruleSet := Set{}
	ruleSet.Rules = []Rule{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ruleSet.Can(tc.role, tc.method, tc.route, tc.resource)
			if err != nil && err.Error() != tc.expError.Error() {
				t.Error(errors.New(fmt.Sprintf("Expected error: %v, got: %v", tc.expError, err)))
			}
		})
	}
}

func TestCannot(t *testing.T) {
	testCases := []struct {
		name         string
		role         utils.Role
		method       utils.HttpMethod
		route        utils.HttpRoute
		resource     utils.Resource
		expStringKey string
		expHashKey   string
		expError     error
	}{
		{
			name:         "Rule should be generated successfully",
			role:         "admin",
			method:       "GET",
			route:        "/books",
			resource:     "book",
			expStringKey: "admin | GET | /books | book",
			expHashKey:   "682af11c5a0f6fe88c9799a215fb41128b34c89c811510ca62d6a8f5c7230592",
			expError:     nil,
		},
		{
			name:         "Duplicate rule should NOT be generated",
			role:         "admin",
			method:       "GET",
			route:        "/books",
			resource:     "book",
			expStringKey: "admin | GET | /books | book",
			expHashKey:   "682af11c5a0f6fe88c9799a215fb41128b34c89c811510ca62d6a8f5c7230592",
			expError:     &utils.ExistingRuleError{StringKey: "admin | GET | /books"},
		},
	}

	ruleSet := Set{}
	ruleSet.Rules = []Rule{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ruleSet.Cannot(tc.role, tc.method, tc.route, tc.resource)
			if err != nil && err.Error() != tc.expError.Error() {
				t.Error(errors.New(fmt.Sprintf("Expected error: %v, got: %v", tc.expError, err)))
			}
		})
	}
}

func TestGenerateCredentials(t *testing.T) {
	type credentials struct {
		role     utils.Role
		resource utils.Resource
		actions  []utils.HttpMethod
	}

	testOneCredentials := credentials{
		role:     "admin",
		resource: "book",
		actions:  []utils.HttpMethod{"GET"},
	}

	testTwoCredentials := credentials{
		role:     "admin",
		resource: "book",
		actions:  []utils.HttpMethod{"GET", "POST"},
	}

	testCases := []struct {
		name               string
		role               utils.Role
		method             utils.HttpMethod
		route              utils.HttpRoute
		resource           utils.Resource
		expStringKey       string
		expHashKey         string
		expError           error
		expCredentials     credentials
		expRuleCount       int
		expCredentialCount int
		expActionCount     int
	}{
		{
			name:               "Admin GET Book rule created",
			role:               "admin",
			method:             "GET",
			route:              "/books",
			resource:           "book",
			expStringKey:       "admin | GET | /books | book",
			expHashKey:         "682af11c5a0f6fe88c9799a215fb41128b34c89c811510ca62d6a8f5c7230592",
			expError:           nil,
			expCredentials:     testOneCredentials,
			expRuleCount:       1,
			expCredentialCount: 1,
			expActionCount:     1,
		},
		{
			name:               "Admin PUT Book rule created",
			role:               "admin",
			method:             "PUT",
			route:              "/books",
			resource:           "book",
			expStringKey:       "admin | PUT | /books | book",
			expHashKey:         "682af11c5a0f6fe88c9799a215fb41128b34c89c811510ca62d6a8f5c7230592",
			expError:           nil,
			expCredentials:     testTwoCredentials,
			expRuleCount:       2,
			expCredentialCount: 1,
			expActionCount:     2,
		},
		{
			name:               "Admin POST Book rule created",
			role:               "admin",
			method:             "POST",
			route:              "/books",
			resource:           "book",
			expStringKey:       "admin | POST | /books | book",
			expHashKey:         "682af11c5a0f6fe88c9799a215fb41128b34c89c811510ca62d6a8f5c7230592",
			expError:           nil,
			expCredentials:     testTwoCredentials,
			expRuleCount:       3,
			expCredentialCount: 1,
			expActionCount:     3,
		},
		{
			name:               "Editor GET Book rule created",
			role:               "editor",
			method:             "GET",
			route:              "/books",
			resource:           "book",
			expStringKey:       "admin | GET | /books | book",
			expHashKey:         "682af11c5a0f6fe88c9799a215fb41128b34c89c811510ca62d6a8f5c7230592",
			expError:           nil,
			expCredentials:     testTwoCredentials,
			expRuleCount:       4,
			expCredentialCount: 2,
			expActionCount:     1,
		},
	}

	ruleSet := Set{}
	ruleSet.Rules = []Rule{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ruleSet.Can(tc.role, tc.method, tc.route, tc.resource)
			if err != nil {
				t.Error(
					errors.New(
						fmt.Sprintf(
							"Expected no error but got: %v",
							err,
						),
					),
				)
			}

			rulesLength := len(ruleSet.Rules)
			if rulesLength != tc.expRuleCount {
				t.Error(
					errors.New(
						fmt.Sprintf(
							"Expected one rule to exist, but found: %v",
							rulesLength,
						),
					),
				)
			}

			credentialsLength := len(ruleSet.Credentials)
			if credentialsLength != tc.expCredentialCount {
				t.Error(
					errors.New(
						fmt.Sprintf(
							"Expected one credential to exist, but found: %v",
							credentialsLength,
						),
					),
				)
			}

			for _, credential := range ruleSet.Credentials {
				if credential.Resource == tc.resource &&
					credential.Role == tc.role &&
					len(credential.Actions) != tc.expActionCount {
					t.Error(
						errors.New(
							fmt.Sprintf(
								"Expected %v action(s) to exist, but found: %v",
								tc.expActionCount,
								len(credential.Actions),
							),
						),
					)
				}
			}
		})
	}
}
