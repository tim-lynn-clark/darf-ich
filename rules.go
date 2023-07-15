package darfich

import "crypto/sha256"

type Rule struct {
	hashKey   string
	stringKey string
	role      Role
	action    Action
	method    HttpMethod
	route     HttpRoute
}

func GenerateRuleKeys(role Role, method HttpMethod, route HttpRoute) (stringKey string, hashKey string) {
	// Build component key from role+method+route
	stringKey = string(role) + " | " + string(method) + " | " + string(route)

	// Create hash of component key
	h := sha256.New()
	h.Write([]byte(stringKey))
	hashKey = string(h.Sum(nil))

	return stringKey, hashKey
}
