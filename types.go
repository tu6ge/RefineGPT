package main

type Assignment struct {
	Agent string `json:"agent"`
	Task  string `json:"task"`
	Time  string `json:"time"`
}

type Decision struct {
	Assignments []Assignment `json:"assignments"`
}

type RuleViolation struct {
	RuleID   string                 `json:"rule_id"`
	Message  string                 `json:"message"`
	Location map[string]interface{} `json:"location"`
}

type Rule interface {
	ID() string
	Check(d Decision) []RuleViolation
}

type LLMClient interface {
	GenerateDecision(feedback []RuleViolation) Decision
}
