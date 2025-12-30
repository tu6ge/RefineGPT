package engine

import "context"

// Generator 是候选解生成器（通常是 LLM）
type Generator interface {
	Generate(ctx context.Context, input GenerateInput) (Candidate, error)
}

type GenerateInput struct {
	State    State
	Feedback []Feedback
}
