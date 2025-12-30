package main

type AgentNotEmptyRule struct{}

func (r AgentNotEmptyRule) ID() string {
	return "AgentNotEmpty"
}

func (r AgentNotEmptyRule) Check(d Decision) []RuleViolation {
	var v []RuleViolation
	for i, a := range d.Assignments {
		if a.Agent == "" {
			v = append(v, RuleViolation{
				RuleID:  r.ID(),
				Message: "agent must not be empty",
				Location: map[string]interface{}{
					"assignment_index": i,
					"field":            "agent",
				},
			})
		}
	}
	return v
}

type NoOverlapRule struct{}

func (r NoOverlapRule) ID() string {
	return "NoOverlap"
}

func (r NoOverlapRule) Check(d Decision) []RuleViolation {
	seen := make(map[string]int)
	var v []RuleViolation

	for i, a := range d.Assignments {
		key := a.Agent + "@" + a.Time
		if j, ok := seen[key]; ok {
			v = append(v, RuleViolation{
				RuleID:  r.ID(),
				Message: "agent has multiple tasks at same time",
				Location: map[string]interface{}{
					"assignment_index": i,
					"conflict_with":    j,
				},
			})
		} else {
			seen[key] = i
		}
	}
	return v
}
