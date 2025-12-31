package engine

import "context"

// Validator 校验 Candidate，返回结构化反馈
type Validator interface {
	Validate(ctx context.Context, state State, c Candidate) ([]Feedback, error)
}
