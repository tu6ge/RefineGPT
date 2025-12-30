package llm

import "github.com/tu6ge/RefineGPT/framework/core"

type MockPlanner struct{}

func (m *MockPlanner) Generate(ctx core.PlanningContext) (core.Decision, error) {
	// 简单 mock：第一次乱给，第二次修正
	if len(ctx.History) == 0 {
		return core.Decision{
			"assignments": []map[string]any{
				{"person": "Alice", "shift": "Morning"},
				{"person": "Alice", "shift": "Morning"},
			},
		}, nil
	}

	return core.Decision{
		"assignments": []map[string]any{
			{"person": "Alice", "shift": "Morning"},
			{"person": "Bob", "shift": "Morning"},
		},
	}, nil
}
