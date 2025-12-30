package main

import "github.com/tu6ge/RefineGPT/framework/core"

type NoDuplicateShift struct{}

func (n *NoDuplicateShift) Name() string {
	return "NoDuplicateShift"
}

func (n *NoDuplicateShift) Validate(d core.Decision) []core.Violation {
	assignments, ok := d["assignments"].([]map[string]any)
	if !ok {
		return nil
	}

	seen := map[string]bool{}
	var violations []core.Violation

	for i, a := range assignments {
		key := a["person"].(string) + "-" + a["shift"].(string)
		if seen[key] {
			violations = append(violations, core.Violation{
				Constraint: n.Name(),
				Path:       "assignments[" + string(rune(i)) + "]",
				Message:    "duplicate shift assignment",
				Hint:       "assign different person or shift",
			})
		}
		seen[key] = true
	}

	return violations
}
