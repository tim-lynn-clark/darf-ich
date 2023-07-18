package utils

import "fmt"

type ExistingRuleError struct {
	StringKey string
}

func (err *ExistingRuleError) Error() string {
	return fmt.Sprintf("Rule already exists for: %v", err.StringKey)
}
