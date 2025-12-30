package generator

import "github.com/tu6ge/RefineGPT/engine"

type GenerateContext struct {
	State    engine.State
	Feedback []engine.Feedback

	// 可选：JSON Schema / 约束说明
	Schema string
}
