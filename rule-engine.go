package main

type RuleEngine struct {
	rules []Rule
}

func NewRuleEngine(rules ...Rule) *RuleEngine {
	return &RuleEngine{rules: rules}
}

func (re *RuleEngine) Validate(d Decision) []RuleViolation {
	var violations []RuleViolation
	for _, r := range re.rules {
		violations = append(violations, r.Check(d)...)
	}
	return violations
}
