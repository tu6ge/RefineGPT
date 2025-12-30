package core

type ValidationResult struct {
	IsValid    bool
	Violations []Violation
	Score      float64
}

type RuleEngine struct {
	Constraints []Constraint
}

func (r *RuleEngine) Validate(d Decision) ValidationResult {
	var violations []Violation

	for _, c := range r.Constraints {
		v := c.Validate(d)
		if len(v) > 0 {
			violations = append(violations, v...)
		}
	}

	return ValidationResult{
		IsValid:    len(violations) == 0,
		Violations: violations,
		Score:      float64(len(violations)) * -1,
	}
}
