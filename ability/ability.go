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
	}
	rules.newCredential(rule)
	rules.Rules[rule.HashKey] = rule

	return rule, nil
}

func (rules *Set) newCredential(rule Rule) {
	// Find credential
	var credential *utils.Credential
	for _, c := range rules.Credentials {
		if c.Role == rule.Role && c.Resource == rule.Resource {
			credential = &c
			break
		}
	}

	// Create credential if not found
	if credential == nil {
		credential = &utils.Credential{
			Role:     rule.Role,
			Resource: rule.Resource,
		}
		rules.Credentials = append(rules.Credentials, *credential)
	}

	// Add Action to credential
	if !utils.Contains(credential.Actions, rule.Method) {
		credential.Actions = append(credential.Actions, rule.Method)
	}
}
