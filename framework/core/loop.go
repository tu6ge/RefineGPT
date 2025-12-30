package core

import (
	"errors"

	"github.com/tu6ge/RefineGPT/framework/schema"
)

type Planner interface {
	Generate(ctx PlanningContext) (Decision, error)
}

type PlanningContext struct {
	Schema     schema.JSONSchema
	History    []DecisionAttempt
	Violations []Violation
	Context    map[string]any
}

type Solver struct {
	Planner      Planner
	RuleEngine   *RuleEngine
	MaxIteration int
}

func (s *Solver) Solve(p Problem) (Decision, error) {
	ctx := PlanningContext{
		Schema:  p.GetSchema(),
		Context: p.GetContext(),
	}

	for i := 0; i < s.MaxIteration; i++ {
		decision, err := s.Planner.Generate(ctx)
		if err != nil {
			return nil, err
		}

		result := s.RuleEngine.Validate(decision)

		ctx.History = append(ctx.History, DecisionAttempt{
			Decision: decision,
			Result:   result,
		})

		if result.IsValid {
			return decision, nil
		}

		ctx.Violations = result.Violations
	}

	return nil, errors.New("no valid solution found")
}
