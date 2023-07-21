package ability

import (
	"errors"
	"fmt"
	"testing"

	"github.com/tim-lynn-clark/darfich/utils"
)

func TestGenerateRuleKeys(t *testing.T) {
	testCases := []struct {
		name         string
		role         utils.Role
		method       utils.HttpMethod
		route        utils.HttpRoute
		resource     utils.Resource
		expStringKey string
		expHashKey   string
		expError     bool
	}{
		{
			name:         "GET keys should be generated correctly",
			role:         "admin",
			method:       "GET",
			route:        "/books",
			resource:     "book",
			expStringKey: "admin | GET | /books | book",
			expHashKey:   "682af11c5a0f6fe88c9799a215fb41128b34c89c811510ca62d6a8f5c7230592",
			expError:     false,
		},
		{
			name:         "POST keys should be generated correctly",
			role:         "admin",
			method:       "POST",
			route:        "/books",
			resource:     "book",
			expStringKey: "admin | POST | /books | book",
			expHashKey:   "884ad4bef9faf3fec8706b9bd8901a15723799db463e5c9d49fc22598db4f89d",
			expError:     false,
		},
		{
			name:         "POST keys should fail due mismatching Role",
			role:         "editor",
			method:       "POST",
			route:        "/books",
			resource:     "book",
			expStringKey: "admin | POST | /books | book",
			expHashKey:   "884ad4bef9faf3fec8706b9bd8901a15723799db463e5c9d49fc22598db4f89d",
			expError:     true,
		},
		{
			name:         "POST keys should fail due to mismatching Method",
			role:         "admin",
			method:       "PATCH",
			route:        "/books",
			resource:     "book",
			expStringKey: "admin | POST | /books | book",
			expHashKey:   "884ad4bef9faf3fec8706b9bd8901a15723799db463e5c9d49fc22598db4f89d",
			expError:     true,
		},
		{
			name:         "POST keys should fail due to mismatching Route",
			role:         "admin",
			method:       "PATCH",
			route:        "/books/",
			resource:     "book",
			expStringKey: "admin | POST | /books | book",
			expHashKey:   "884ad4bef9faf3fec8706b9bd8901a15723799db463e5c9d49fc22598db4f89d",
			expError:     true,
		},
		{
			name:         "POST keys should fail due to mismatching Resource",
			role:         "admin",
			method:       "PATCH",
			route:        "/books",
			resource:     "document",
			expStringKey: "admin | POST | /books | book",
			expHashKey:   "884ad4bef9faf3fec8706b9bd8901a15723799db463e5c9d49fc22598db4f89d",
			expError:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stringKey, hashKey := GenerateRuleKeys(tc.role, tc.method, tc.route)

			var err error

			if stringKey != tc.expStringKey {
				err = errors.New(fmt.Sprintf("Expected StringKey to be %v, got %v", tc.expStringKey, stringKey))
			}

			if hashKey != tc.expHashKey {
				err = errors.New(fmt.Sprintf("Expected HashKey to be %v, got %v", tc.expHashKey, hashKey))
			}

			if tc.expError && err == nil {
				t.Errorf(err.Error())
			}
		})
	}
}
