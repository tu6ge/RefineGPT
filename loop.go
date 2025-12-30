package main

import "fmt"

type LoopManager struct {
	MaxIterations int
	LLM           LLMClient
	Engine        *RuleEngine
}

func (lm *LoopManager) Run() (Decision, error) {
	var feedback []RuleViolation

	for i := 0; i < lm.MaxIterations; i++ {
		decision := lm.LLM.GenerateDecision(feedback)
		violations := lm.Engine.Validate(decision)

		if len(violations) == 0 {
			return decision, nil
		}

		feedback = violations
		fmt.Printf("Iteration %d violations:\n%+v\n\n", i+1, violations)
	}

	return Decision{}, fmt.Errorf("no valid decision found")
}
