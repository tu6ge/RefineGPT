package engine

import (
	"context"
)

// Candidate 表示 LLM 生成的候选解
// 框架只关心 JSON，不关心结构
type Candidate interface {
	Raw() []byte
	As(v any) error
}

type CandidateFactory interface {
	FromLLMOutput(
		ctx context.Context,
		raw string,
	) (Candidate, error)
}
