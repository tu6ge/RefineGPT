package core

// 限制
type Constraint interface {
	Name() string
	Validate(decision Decision) []Violation
}

type Violation struct {
	Constraint string
	Path       string
	Message    string
	Hint       string
}
