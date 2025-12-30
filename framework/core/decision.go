package core

// 决策
type Decision map[string]any

// 决策尝试
type DecisionAttempt struct {
	Decision Decision
	Result   ValidationResult
}
