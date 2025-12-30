package llm

import "github.com/tu6ge/RefineGPT/framework/core"

type Planner interface {
	Generate(ctx core.PlanningContext) (core.Decision, error)
}
