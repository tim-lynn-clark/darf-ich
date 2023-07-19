package ability

import (
	"github.com/tim-lynn-clark/darfich/utils"
)

type Set struct {
	Rules       map[string]Rule    `json:"rules"`
	Credentials []utils.Credential `json:"credentials"`
}

func (rules *Set) Can(role utils.Role, method utils.HttpMethod,
	route utils.HttpRoute, resource utils.Resource) (Rule, error) {
	return rules.newRule(utils.ActionAllow, role, method, route, resource)
}

func (rules *Set) Cannot(role utils.Role, method utils.HttpMethod,
	route utils.HttpRoute, resource utils.Resource) (Rule, error) {
	return rules.newRule(utils.ActionAllow, role, method, route, resource)
}

func (rules *Set) newRule(action utils.Action, role utils.Role,
	method utils.HttpMethod, route utils.HttpRoute, resource utils.Resource) (Rule, error) {
	stringKey, hashKey := GenerateRuleKeys(role, method, route, resource)

	// Verify the rule doesn't already exist
	_, ok := rules.Rules[hashKey]
	if ok {
		return Rule{}, &utils.ExistingRuleError{
			StringKey: stringKey,
		}
	}

	// Create rule
	rule := Rule{
		HashKey:   hashKey,
		StringKey: stringKey,
		Role:      role,
		Action:    action,
		Method:    method,
		Route:     route,
		Resource:  resource,
	}
	rules.newCredential(rule)
	rules.Rules[rule.HashKey] = rule

	return rule, nil
}

func (rules *Set) newCredential(rule Rule) {
	// Find credential
	for idx, c := range rules.Credentials {
		if c.Role == rule.Role && c.Resource == rule.Resource {
			// Add Action to existing credential
			if !utils.Contains(c.Actions, rule.Method) {
				rules.Credentials[idx].Actions = append(c.Actions, rule.Method)
			}
			// Short circuit since credential already exists
			return
		}
	}

	// Create credential
	credential := &utils.Credential{
		Role:     rule.Role,
		Resource: rule.Resource,
		Actions:  []utils.HttpMethod{},
	}
	credential.Actions = append(credential.Actions, rule.Method)
	rules.Credentials = append(rules.Credentials, *credential)
}
