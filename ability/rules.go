package ability

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/tim-lynn-clark/darfich/utils"
)

type Rule struct {
	HashKey   string
	StringKey string
	Role      utils.Role
	Action    utils.Action
	Method    utils.HttpMethod
	Route     utils.HttpRoute
	Resource  utils.Resource
}

func GenerateRuleKeys(role utils.Role, method utils.HttpMethod, route utils.HttpRoute, resource utils.Resource) (stringKey string, hashKey string) {
	// Build component key from Role+Method+Route
	stringKey = string(role) + " | " + string(method) + " | " + string(route) + " | " + string(resource)

	// Create hash of component key
	sum := sha256.Sum256([]byte(stringKey))
	hashKey = hex.EncodeToString(sum[:])

	return stringKey, hashKey
}
