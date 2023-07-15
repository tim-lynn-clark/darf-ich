package darfich

import (
	"log"
)

type Abilities struct {
	rules map[string]Rule
}

func (abilities *Abilities) Can(role Role, action Action, method HttpMethod, route HttpRoute) {
	stringKey, hashKey := GenerateRuleKeys(role, method, route)

	// Verify the rule doesn't already exist
	_, ok := abilities.rules[hashKey]
	if ok {
		log.Print("Rule already exists")
		return
	}

	// Create rule
	rule := Rule{
		hashKey:   hashKey,
		stringKey: stringKey,
		role:      role,
		action:    action,
		method:    method,
		route:     route,
	}

	abilities.rules[rule.hashKey] = rule
}
